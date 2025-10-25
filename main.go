// main.go - Plugin entry point
package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
	"github.com/RuneORakeie/tflint-ruleset-fabric/rules"
	"github.com/RuneORakeie/tflint-ruleset-fabric/rules/apispec"
)

var (
	Sha1ver = ""
)

func main() {
	// Combine custom rules with generated rules
	allRules := []tflint.Rule{
		// Workspace rules
		rules.NewFabricWorkspaceCapacity(),
		rules.NewFabricWorkspaceDescription(),
		rules.NewFabricWorkspaceRoleAssignmentRole(),
		
		// Role assignment rules
		rules.NewFabricRoleAssignmentRecommended(),
		
		// Capacity rules
		rules.NewFabricCapacityRegion(),
		
		// Item rules
		rules.NewFabricItemDescriptionRecommended(),
		
		// Deployment pipeline rules
		rules.NewFabricDeploymentPipelineStagesCount(),
		rules.NewFabricDeploymentPipelineStagesDisplayNameLength(),
		rules.NewFabricDeploymentPipelineStagesDescriptionLength(),
		
		// Domain rules
		rules.NewFabricDomainContributorsScope(),
		
		// Git integration validation rules
		rules.NewFabricWorkspaceGitProviderType(),
		rules.NewFabricWorkspaceGitInitializationStrategy(),
		rules.NewFabricWorkspaceGitDirectoryName(),
		rules.NewFabricWorkspaceGitCredentialsSource(),
		rules.NewFabricWorkspaceGitAzureDevOpsAttributes(),
		rules.NewFabricWorkspaceGitGitHubAttributes(),
		rules.NewFabricWorkspaceGitStringLengths(),
	}
	
	// Add all auto-generated rules from apispec
	allRules = append(allRules, apispec.Rules()...)
	
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "fabric",
			Version: project.Version,
			Rules:   allRules,
		},
	})
}
