package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricWorkspaceGitAzureDevOpsAttributes validates required attributes for AzureDevOps provider
type FabricWorkspaceGitAzureDevOpsAttributes struct {
	tflint.DefaultRule
}

func NewFabricWorkspaceGitAzureDevOpsAttributes() *FabricWorkspaceGitAzureDevOpsAttributes {
	return &FabricWorkspaceGitAzureDevOpsAttributes{}
}

func (r *FabricWorkspaceGitAzureDevOpsAttributes) Name() string {
	return "fabric_workspace_git_azdo_attributes_required"
}

func (r *FabricWorkspaceGitAzureDevOpsAttributes) Enabled() bool {
	return true
}

func (r *FabricWorkspaceGitAzureDevOpsAttributes) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *FabricWorkspaceGitAzureDevOpsAttributes) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *FabricWorkspaceGitAzureDevOpsAttributes) Check(runner tflint.Runner) error {
	resourceContent, err := runner.GetResourceContent("fabric_workspace_git", &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "git_provider_details",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "git_provider_type"},
						{Name: "organization_name"},
						{Name: "project_name"},
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
				runner.EvaluateExpr(attr.Expr, &providerType, nil)
			}
			
			// Only validate if provider is AzureDevOps
			if providerType == "AzureDevOps" {
				// Check organization_name
				if attr, exists := block.Body.Attributes["organization_name"]; !exists || attr.Expr == nil {
					runner.EmitIssue(
						r,
						"organization_name is required when git_provider_type is 'AzureDevOps'",
						block.DefRange,
					)
				}
				
				// Check project_name
				if attr, exists := block.Body.Attributes["project_name"]; !exists || attr.Expr == nil {
					runner.EmitIssue(
						r,
						"project_name is required when git_provider_type is 'AzureDevOps'",
						block.DefRange,
					)
				}
			}
		}
	}

	return nil
}
