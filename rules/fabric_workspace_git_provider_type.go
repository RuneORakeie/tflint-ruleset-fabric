package rules

import (
	"fmt"

	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// FabricWorkspaceGitProviderType validates Git provider type
type FabricWorkspaceGitProviderType struct {
	tflint.DefaultRule
}

func NewFabricWorkspaceGitProviderType() *FabricWorkspaceGitProviderType {
	return &FabricWorkspaceGitProviderType{}
}

func (r *FabricWorkspaceGitProviderType) Name() string {
	return "fabric_workspace_git_provider_type_valid"
}

func (r *FabricWorkspaceGitProviderType) Enabled() bool {
	return true
}

func (r *FabricWorkspaceGitProviderType) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *FabricWorkspaceGitProviderType) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *FabricWorkspaceGitProviderType) Check(runner tflint.Runner) error {
	resourceContent, err := runner.GetResourceContent("fabric_workspace_git", &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "git_provider_details",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "git_provider_type"},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	validProviders := map[string]bool{
		"AzureDevOps": true,
		"GitHub":      true,
	}

	for _, resource := range resourceContent.Blocks {
		gitProviderBlocks := resource.Body.Blocks.OfType("git_provider_details")

		for _, block := range gitProviderBlocks {
			if attr, exists := block.Body.Attributes["git_provider_type"]; exists && attr.Expr != nil {
				var provider string
				if err := runner.EvaluateExpr(attr.Expr, &provider, nil); err == nil && provider != "" {
					if !validProviders[provider] {
						runner.EmitIssue(
							r,
							fmt.Sprintf("Invalid git_provider_type '%s'. Must be one of: AzureDevOps, GitHub", provider),
							attr.Range,
						)
					}
				}
			}
		}
	}

	return nil
}
