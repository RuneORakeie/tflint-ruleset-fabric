package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricWorkspaceGitCredentialsSource validates git_credentials.source based on provider type
type FabricWorkspaceGitCredentialsSource struct {
	tflint.DefaultRule
}

func NewFabricWorkspaceGitCredentialsSource() *FabricWorkspaceGitCredentialsSource {
	return &FabricWorkspaceGitCredentialsSource{}
}

func (r *FabricWorkspaceGitCredentialsSource) Name() string {
	return "fabric_workspace_git_credentials_source_valid"
}

func (r *FabricWorkspaceGitCredentialsSource) Enabled() bool {
	return true
}

func (r *FabricWorkspaceGitCredentialsSource) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *FabricWorkspaceGitCredentialsSource) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *FabricWorkspaceGitCredentialsSource) Check(runner tflint.Runner) error {
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
			{
				Type: "git_credentials",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "source"},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resourceContent.Blocks {
		var providerType string
		
		// Get provider type
		gitProviderBlocks := resource.Body.Blocks.OfType("git_provider_details")
		for _, block := range gitProviderBlocks {
			if attr, exists := block.Body.Attributes["git_provider_type"]; exists && attr.Expr != nil {
				runner.EvaluateExpr(attr.Expr, &providerType, nil)
			}
		}
		
		// Check credentials source
		gitCredentialsBlocks := resource.Body.Blocks.OfType("git_credentials")
		for _, block := range gitCredentialsBlocks {
			if attr, exists := block.Body.Attributes["source"]; exists && attr.Expr != nil {
				var source string
				if err := runner.EvaluateExpr(attr.Expr, &source, nil); err == nil && source != "" {
					// GitHub only supports ConfiguredConnection
					if providerType == "GitHub" && source != "ConfiguredConnection" {
						runner.EmitIssue(
							r,
							fmt.Sprintf("GitHub only supports git_credentials.source = 'ConfiguredConnection' (current: '%s')", source),
							attr.Range,
						)
					}
					
					// AzureDevOps supports ConfiguredConnection or Automatic
					if providerType == "AzureDevOps" && source != "ConfiguredConnection" && source != "Automatic" {
						runner.EmitIssue(
							r,
							fmt.Sprintf("Azure DevOps git_credentials.source must be 'ConfiguredConnection' or 'Automatic' (current: '%s')", source),
							attr.Range,
						)
					}
				}
			}
		}
	}

	return nil
}
