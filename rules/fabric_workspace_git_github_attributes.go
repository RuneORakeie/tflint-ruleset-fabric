package rules

import (
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// FabricWorkspaceGitGitHubAttributes validates required attributes for GitHub provider
type FabricWorkspaceGitGitHubAttributes struct {
	tflint.DefaultRule
}

func NewFabricWorkspaceGitGitHubAttributes() *FabricWorkspaceGitGitHubAttributes {
	return &FabricWorkspaceGitGitHubAttributes{}
}

func (r *FabricWorkspaceGitGitHubAttributes) Name() string {
	return "fabric_workspace_git_github_attributes_required"
}

func (r *FabricWorkspaceGitGitHubAttributes) Enabled() bool {
	return true
}

func (r *FabricWorkspaceGitGitHubAttributes) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *FabricWorkspaceGitGitHubAttributes) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *FabricWorkspaceGitGitHubAttributes) Check(runner tflint.Runner) error {
	resourceContent, err := runner.GetResourceContent("fabric_workspace_git", &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "git_provider_details",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "git_provider_type"},
						{Name: "owner_name"},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resourceContent.Blocks {
		gitProviderBlocks := resource.Body.Blocks.OfType("git_provider_details")

		for _, block := range gitProviderBlocks {
			var providerType string
			if attr, exists := block.Body.Attributes["git_provider_type"]; exists && attr.Expr != nil {
				_ = runner.EvaluateExpr(attr.Expr, &providerType, nil)
			}

			// Only validate if provider is GitHub
			if providerType == "GitHub" {
				// Check owner_name
				if attr, exists := block.Body.Attributes["owner_name"]; !exists || attr.Expr == nil {
					runner.EmitIssue(
						r,
						"owner_name is required when git_provider_type is 'GitHub'",
						block.DefRange,
					)
				}
			}
		}
	}

	return nil
}
