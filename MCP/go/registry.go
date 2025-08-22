package main

import (
	"github.com/aws-resource-groups/mcp-server/config"
	"github.com/aws-resource-groups/mcp-server/models"
	tools_update_group "github.com/aws-resource-groups/mcp-server/tools/update_group"
	tools_list_group_resources "github.com/aws-resource-groups/mcp-server/tools/list_group_resources"
	tools_get_group_query "github.com/aws-resource-groups/mcp-server/tools/get_group_query"
	tools_put_group_configuration "github.com/aws-resource-groups/mcp-server/tools/put_group_configuration"
	tools_update_group_query "github.com/aws-resource-groups/mcp-server/tools/update_group_query"
	tools_get_group "github.com/aws-resource-groups/mcp-server/tools/get_group"
	tools_get_account_settings "github.com/aws-resource-groups/mcp-server/tools/get_account_settings"
	tools_get_group_configuration "github.com/aws-resource-groups/mcp-server/tools/get_group_configuration"
	tools_groups_list "github.com/aws-resource-groups/mcp-server/tools/groups_list"
	tools_resources "github.com/aws-resource-groups/mcp-server/tools/resources"
	tools_ungroup_resources "github.com/aws-resource-groups/mcp-server/tools/ungroup_resources"
	tools_group_resources "github.com/aws-resource-groups/mcp-server/tools/group_resources"
	tools_delete_group "github.com/aws-resource-groups/mcp-server/tools/delete_group"
	tools_update_account_settings "github.com/aws-resource-groups/mcp-server/tools/update_account_settings"
	tools_groups "github.com/aws-resource-groups/mcp-server/tools/groups"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_update_group.CreateUpdategroupTool(cfg),
		tools_list_group_resources.CreateListgroupresourcesTool(cfg),
		tools_get_group_query.CreateGetgroupqueryTool(cfg),
		tools_put_group_configuration.CreatePutgroupconfigurationTool(cfg),
		tools_update_group_query.CreateUpdategroupqueryTool(cfg),
		tools_get_group.CreateGetgroupTool(cfg),
		tools_get_account_settings.CreateGetaccountsettingsTool(cfg),
		tools_get_group_configuration.CreateGetgroupconfigurationTool(cfg),
		tools_groups_list.CreateListgroupsTool(cfg),
		tools_resources.CreateUntagTool(cfg),
		tools_resources.CreateTagTool(cfg),
		tools_resources.CreateGettagsTool(cfg),
		tools_ungroup_resources.CreateUngroupresourcesTool(cfg),
		tools_group_resources.CreateGroupresourcesTool(cfg),
		tools_resources.CreateSearchresourcesTool(cfg),
		tools_delete_group.CreateDeletegroupTool(cfg),
		tools_update_account_settings.CreateUpdateaccountsettingsTool(cfg),
		tools_groups.CreateCreategroupTool(cfg),
	}
}
