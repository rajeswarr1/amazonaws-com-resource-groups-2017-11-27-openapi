package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/aws-resource-groups/mcp-server/config"
	"github.com/aws-resource-groups/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GettagsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		ArnVal, ok := args["Arn"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: Arn"), nil
		}
		Arn, ok := ArnVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: Arn"), nil
		}
		url := fmt.Sprintf("%s/resources/%s/tags", cfg.BaseURL, Arn)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.GetTagsOutput
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGettagsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_resources_Arn_tags",
		mcp.WithDescription("<p>Returns a list of tags that are associated with a resource group, specified by an ARN.</p> <p> <b>Minimum permissions</b> </p> <p>To run this command, you must have the following permissions:</p> <ul> <li> <p> <code>resource-groups:GetTags</code> </p> </li> </ul>"),
		mcp.WithString("Arn", mcp.Required(), mcp.Description("The ARN of the resource group whose tags you want to retrieve.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GettagsHandler(cfg),
	}
}
