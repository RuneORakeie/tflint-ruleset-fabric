package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricWorkspaceGitStringLengths validates string length constraints for git_provider_details attributes
type FabricWorkspaceGitStringLengths struct {
	tflint.DefaultRule
}

func NewFabricWorkspaceGitStringLengths() *FabricWorkspaceGitStringLengths {
	return &FabricWorkspaceGitStringLengths{}
}

func (r *FabricWorkspaceGitStringLengths) Name() string {
	return "fabric_workspace_git_string_lengths"
}

func (r *FabricWorkspaceGitStringLengths) Enabled() bool {
	return true
}

func (r *FabricWorkspaceGitStringLengths) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *FabricWorkspaceGitStringLengths) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *FabricWorkspaceGitStringLengths) Check(runner tflint.Runner) error {
	resourceContent, err := runner.GetResourceContent("fabric_workspace_git", &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "git_provider_details",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "branch_name"},
						{Name: "repository_name"},
						{Name: "organization_name"},
						{Name: "owner_name"},
						{Name: "project_name"},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	// Define max lengths for each attribute
	maxLengths := map[string]int{
		"branch_name":       250,
		"repository_name":   128,
		"organization_name": 100,
		"owner_name":        100,
		"project_name":      100,
	}

	for _, resource := range resourceContent.Blocks {
		gitProviderBlocks := resource.Body.Blocks.OfType("git_provider_details")

		for _, block := range gitProviderBlocks {
			// Check each attribute
			for attrName, maxLength := range maxLengths {
				if attr, exists := block.Body.Attributes[attrName]; exists && attr.Expr != nil {
					var value string
					if err := runner.EvaluateExpr(attr.Expr, &value, nil); err == nil && value != "" {
						if len(value) > maxLength {
							if err := runner.EmitIssue(
								r,
								fmt.Sprintf("%s must not exceed %d characters (current: %d)", attrName, maxLength, len(value)),
								attr.Range,
							); err != nil {
								return err
							}
						}
					}
				}
			}
		}
	}

	return nil
}
