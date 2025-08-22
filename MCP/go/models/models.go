package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GroupResourcesInput represents the GroupResourcesInput schema from the OpenAPI specification
type GroupResourcesInput struct {
	Resourcearns interface{} `json:"ResourceArns"`
	Group interface{} `json:"Group"`
}

// UpdateGroupInput represents the UpdateGroupInput schema from the OpenAPI specification
type UpdateGroupInput struct {
	Description interface{} `json:"Description,omitempty"`
	Group interface{} `json:"Group,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
}

// GroupConfiguration represents the GroupConfiguration schema from the OpenAPI specification
type GroupConfiguration struct {
	Proposedconfiguration interface{} `json:"ProposedConfiguration,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Configuration interface{} `json:"Configuration,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
}

// Group represents the Group schema from the OpenAPI specification
type Group struct {
	Description interface{} `json:"Description,omitempty"`
	Grouparn interface{} `json:"GroupArn"`
	Name interface{} `json:"Name"`
}

// UpdateAccountSettingsOutput represents the UpdateAccountSettingsOutput schema from the OpenAPI specification
type UpdateAccountSettingsOutput struct {
	Accountsettings interface{} `json:"AccountSettings,omitempty"`
}

// GroupConfigurationParameter represents the GroupConfigurationParameter schema from the OpenAPI specification
type GroupConfigurationParameter struct {
	Name interface{} `json:"Name"`
	Values interface{} `json:"Values,omitempty"`
}

// ListGroupsOutput represents the ListGroupsOutput schema from the OpenAPI specification
type ListGroupsOutput struct {
	Groupidentifiers interface{} `json:"GroupIdentifiers,omitempty"`
	Groups interface{} `json:"Groups,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// QueryError represents the QueryError schema from the OpenAPI specification
type QueryError struct {
	Errorcode interface{} `json:"ErrorCode,omitempty"`
	Message interface{} `json:"Message,omitempty"`
}

// UpdateAccountSettingsInput represents the UpdateAccountSettingsInput schema from the OpenAPI specification
type UpdateAccountSettingsInput struct {
	Grouplifecycleeventsdesiredstatus interface{} `json:"GroupLifecycleEventsDesiredStatus,omitempty"`
}

// UngroupResourcesOutput represents the UngroupResourcesOutput schema from the OpenAPI specification
type UngroupResourcesOutput struct {
	Failed interface{} `json:"Failed,omitempty"`
	Pending interface{} `json:"Pending,omitempty"`
	Succeeded interface{} `json:"Succeeded,omitempty"`
}

// GetGroupConfigurationInput represents the GetGroupConfigurationInput schema from the OpenAPI specification
type GetGroupConfigurationInput struct {
	Group interface{} `json:"Group,omitempty"`
}

// CreateGroupOutput represents the CreateGroupOutput schema from the OpenAPI specification
type CreateGroupOutput struct {
	Group interface{} `json:"Group,omitempty"`
	Groupconfiguration interface{} `json:"GroupConfiguration,omitempty"`
	Resourcequery interface{} `json:"ResourceQuery,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// PendingResource represents the PendingResource schema from the OpenAPI specification
type PendingResource struct {
	Resourcearn interface{} `json:"ResourceArn,omitempty"`
}

// ListGroupResourcesItem represents the ListGroupResourcesItem schema from the OpenAPI specification
type ListGroupResourcesItem struct {
	Status interface{} `json:"Status,omitempty"`
	Identifier ResourceIdentifier `json:"Identifier,omitempty"` // A structure that contains the ARN of a resource and its resource type.
}

// Tags represents the Tags schema from the OpenAPI specification
type Tags struct {
}

// GetTagsOutput represents the GetTagsOutput schema from the OpenAPI specification
type GetTagsOutput struct {
	Arn interface{} `json:"Arn,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ListGroupsInput represents the ListGroupsInput schema from the OpenAPI specification
type ListGroupsInput struct {
	Filters interface{} `json:"Filters,omitempty"`
}

// UpdateGroupQueryInput represents the UpdateGroupQueryInput schema from the OpenAPI specification
type UpdateGroupQueryInput struct {
	Group interface{} `json:"Group,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
	Resourcequery interface{} `json:"ResourceQuery"`
}

// ResourceQuery represents the ResourceQuery schema from the OpenAPI specification
type ResourceQuery struct {
	Query interface{} `json:"Query"`
	TypeField interface{} `json:"Type"`
}

// UpdateGroupQueryOutput represents the UpdateGroupQueryOutput schema from the OpenAPI specification
type UpdateGroupQueryOutput struct {
	Groupquery interface{} `json:"GroupQuery,omitempty"`
}

// GetGroupOutput represents the GetGroupOutput schema from the OpenAPI specification
type GetGroupOutput struct {
	Group interface{} `json:"Group,omitempty"`
}

// UntagOutput represents the UntagOutput schema from the OpenAPI specification
type UntagOutput struct {
	Keys interface{} `json:"Keys,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
}

// CreateGroupInput represents the CreateGroupInput schema from the OpenAPI specification
type CreateGroupInput struct {
	Configuration interface{} `json:"Configuration,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name"`
	Resourcequery interface{} `json:"ResourceQuery,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// SearchResourcesInput represents the SearchResourcesInput schema from the OpenAPI specification
type SearchResourcesInput struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resourcequery interface{} `json:"ResourceQuery"`
}

// GroupIdentifier represents the GroupIdentifier schema from the OpenAPI specification
type GroupIdentifier struct {
	Groupname interface{} `json:"GroupName,omitempty"`
	Grouparn interface{} `json:"GroupArn,omitempty"`
}

// TagInput represents the TagInput schema from the OpenAPI specification
type TagInput struct {
	Tags interface{} `json:"Tags"`
}

// GetGroupInput represents the GetGroupInput schema from the OpenAPI specification
type GetGroupInput struct {
	Group interface{} `json:"Group,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
}

// GetGroupConfigurationOutput represents the GetGroupConfigurationOutput schema from the OpenAPI specification
type GetGroupConfigurationOutput struct {
	Groupconfiguration interface{} `json:"GroupConfiguration,omitempty"`
}

// ResourceStatus represents the ResourceStatus schema from the OpenAPI specification
type ResourceStatus struct {
	Name interface{} `json:"Name,omitempty"`
}

// PutGroupConfigurationInput represents the PutGroupConfigurationInput schema from the OpenAPI specification
type PutGroupConfigurationInput struct {
	Configuration interface{} `json:"Configuration,omitempty"`
	Group interface{} `json:"Group,omitempty"`
}

// GetTagsInput represents the GetTagsInput schema from the OpenAPI specification
type GetTagsInput struct {
}

// SearchResourcesOutput represents the SearchResourcesOutput schema from the OpenAPI specification
type SearchResourcesOutput struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Queryerrors interface{} `json:"QueryErrors,omitempty"`
	Resourceidentifiers interface{} `json:"ResourceIdentifiers,omitempty"`
}

// AccountSettings represents the AccountSettings schema from the OpenAPI specification
type AccountSettings struct {
	Grouplifecycleeventsstatus interface{} `json:"GroupLifecycleEventsStatus,omitempty"`
	Grouplifecycleeventsstatusmessage interface{} `json:"GroupLifecycleEventsStatusMessage,omitempty"`
	Grouplifecycleeventsdesiredstatus interface{} `json:"GroupLifecycleEventsDesiredStatus,omitempty"`
}

// TagOutput represents the TagOutput schema from the OpenAPI specification
type TagOutput struct {
	Arn interface{} `json:"Arn,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ResourceFilter represents the ResourceFilter schema from the OpenAPI specification
type ResourceFilter struct {
	Name interface{} `json:"Name"`
	Values interface{} `json:"Values"`
}

// GroupQuery represents the GroupQuery schema from the OpenAPI specification
type GroupQuery struct {
	Groupname interface{} `json:"GroupName"`
	Resourcequery interface{} `json:"ResourceQuery"`
}

// GroupResourcesOutput represents the GroupResourcesOutput schema from the OpenAPI specification
type GroupResourcesOutput struct {
	Failed interface{} `json:"Failed,omitempty"`
	Pending interface{} `json:"Pending,omitempty"`
	Succeeded interface{} `json:"Succeeded,omitempty"`
}

// GetGroupQueryInput represents the GetGroupQueryInput schema from the OpenAPI specification
type GetGroupQueryInput struct {
	Group interface{} `json:"Group,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
}

// ListGroupResourcesOutput represents the ListGroupResourcesOutput schema from the OpenAPI specification
type ListGroupResourcesOutput struct {
	Resourceidentifiers interface{} `json:"ResourceIdentifiers,omitempty"`
	Resources interface{} `json:"Resources,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Queryerrors interface{} `json:"QueryErrors,omitempty"`
}

// PutGroupConfigurationOutput represents the PutGroupConfigurationOutput schema from the OpenAPI specification
type PutGroupConfigurationOutput struct {
}

// UntagInput represents the UntagInput schema from the OpenAPI specification
type UntagInput struct {
	Keys interface{} `json:"Keys"`
}

// ListGroupResourcesInput represents the ListGroupResourcesInput schema from the OpenAPI specification
type ListGroupResourcesInput struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
	Group interface{} `json:"Group,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
}

// GroupFilter represents the GroupFilter schema from the OpenAPI specification
type GroupFilter struct {
	Name interface{} `json:"Name"`
	Values interface{} `json:"Values"`
}

// DeleteGroupOutput represents the DeleteGroupOutput schema from the OpenAPI specification
type DeleteGroupOutput struct {
	Group interface{} `json:"Group,omitempty"`
}

// FailedResource represents the FailedResource schema from the OpenAPI specification
type FailedResource struct {
	Resourcearn interface{} `json:"ResourceArn,omitempty"`
	Errorcode interface{} `json:"ErrorCode,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
}

// UngroupResourcesInput represents the UngroupResourcesInput schema from the OpenAPI specification
type UngroupResourcesInput struct {
	Group interface{} `json:"Group"`
	Resourcearns interface{} `json:"ResourceArns"`
}

// UpdateGroupOutput represents the UpdateGroupOutput schema from the OpenAPI specification
type UpdateGroupOutput struct {
	Group interface{} `json:"Group,omitempty"`
}

// GetAccountSettingsOutput represents the GetAccountSettingsOutput schema from the OpenAPI specification
type GetAccountSettingsOutput struct {
	Accountsettings interface{} `json:"AccountSettings,omitempty"`
}

// GetGroupQueryOutput represents the GetGroupQueryOutput schema from the OpenAPI specification
type GetGroupQueryOutput struct {
	Groupquery interface{} `json:"GroupQuery,omitempty"`
}

// DeleteGroupInput represents the DeleteGroupInput schema from the OpenAPI specification
type DeleteGroupInput struct {
	Group interface{} `json:"Group,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
}

// GroupConfigurationItem represents the GroupConfigurationItem schema from the OpenAPI specification
type GroupConfigurationItem struct {
	Parameters interface{} `json:"Parameters,omitempty"`
	TypeField interface{} `json:"Type"`
}

// ResourceIdentifier represents the ResourceIdentifier schema from the OpenAPI specification
type ResourceIdentifier struct {
	Resourcearn interface{} `json:"ResourceArn,omitempty"`
	Resourcetype interface{} `json:"ResourceType,omitempty"`
}
