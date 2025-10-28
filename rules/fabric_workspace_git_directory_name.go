package rules

import (
	"fmt"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricWorkspaceGitDirectoryName validates directory_name format and length
type FabricWorkspaceGitDirectoryName struct {
	tflint.DefaultRule
}

func NewFabricWorkspaceGitDirectoryName() *FabricWorkspaceGitDirectoryName {
	return &FabricWorkspaceGitDirectoryName{}
}

func (r *FabricWorkspaceGitDirectoryName) Name() string {
	return "fabric_workspace_git_directory_name_format"
}

func (r *FabricWorkspaceGitDirectoryName) Enabled() bool {
	return true
}

func (r *FabricWorkspaceGitDirectoryName) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *FabricWorkspaceGitDirectoryName) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *FabricWorkspaceGitDirectoryName) Check(runner tflint.Runner) error {
	resourceContent, err := runner.GetResourceContent("fabric_workspace_git", &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "git_provider_details",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "directory_name"},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	const maxLength = 256

	for _, resource := range resourceContent.Blocks {
		gitProviderBlocks := resource.Body.Blocks.OfType("git_provider_details")

		for _, block := range gitProviderBlocks {
			if attr, exists := block.Body.Attributes["directory_name"]; exists && attr.Expr != nil {
				var directoryName string
				if err := runner.EvaluateExpr(attr.Expr, &directoryName, nil); err == nil && directoryName != "" {
					// Check if starts with /
					if !strings.HasPrefix(directoryName, "/") {
						if err := runner.EmitIssue(
							r,
							"directory_name must start with forward slash '/'",
							attr.Range,
						); err != nil {
							return err
						}
					}

					// Check length
					if len(directoryName) > maxLength {
						if err := runner.EmitIssue(
							r,
							fmt.Sprintf("directory_name must not exceed %d characters (current: %d)", maxLength, len(directoryName)),
							attr.Range,
						); err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}
