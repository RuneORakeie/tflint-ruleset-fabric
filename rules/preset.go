package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/RuneORakeie/tflint-ruleset-fabric/rules/apispec"
)

// Define rule groups that compose presets
var (
	// minimalRules contains only critical validation rules
	minimalRules = []tflint.Rule{
		NewFabricWorkspaceCapacity(),
		NewFabricCapacityRegion(),
	}

	// recommendedOnlyRules are additional rules for the recommended preset
	recommendedOnlyRules = []tflint.Rule{
		NewFabricWorkspaceRoleAssignmentRole(),
		NewFabricDeploymentPipelineStagesCount(),
		NewFabricDomainContributorsScope(),
	}

	// allOnlyRules are additional rules for the all preset
	allOnlyRules = []tflint.Rule{
		NewFabricItemDescriptionRecommended(),
		NewFabricRoleAssignmentRecommended(),
		NewFabricDeploymentPipelineStagesDescriptionLength(),
		NewFabricDeploymentPipelineStagesDisplayNameLength(),
		// Git integration validation rules
		NewFabricWorkspaceGitProviderType(),
		NewFabricWorkspaceGitInitializationStrategy(),
		NewFabricWorkspaceGitDirectoryName(),
		NewFabricWorkspaceGitCredentialsSource(),
		NewFabricWorkspaceGitAzureDevOpsAttributes(),
		NewFabricWorkspaceGitGitHubAttributes(),
		NewFabricWorkspaceGitStringLengths(),
	}
)

// PresetRules maps preset names to their rule sets
// Presets are composed hierarchically:
// - minimal: Core validation rules
// - recommended: minimal + common best practices
// - all: recommended + all business logic rules + all generated rules
var PresetRules = map[string][]tflint.Rule{
	"minimal": minimalRules,

	"recommended": append(
		minimalRules,
		recommendedOnlyRules...,
	),

	"all": append(
		append(
			append(minimalRules, recommendedOnlyRules...),
			allOnlyRules...,
		),
		apispec.Rules()...,
	),
}
