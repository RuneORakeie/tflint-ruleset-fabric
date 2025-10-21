package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// FabricWorkspaceDescription ensures workspaces have descriptions
type FabricWorkspaceDescription struct {
	tflint.DefaultRule
}

func NewFabricWorkspaceDescription() *FabricWorkspaceDescription {
	return &FabricWorkspaceDescription{}
}

func (r *FabricWorkspaceDescription) Name() string {
	return "fabric_workspace_description_required"
}

func (r *FabricWorkspaceDescription) Enabled() bool {
	return true
}

func (r *FabricWorkspaceDescription) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *FabricWorkspaceDescription) Link() string {
	return "https://learn.microsoft.com/en-us/fabric/admin/fabric-governance"
}

func (r *FabricWorkspaceDescription) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("fabric_workspace", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "description"},
		},
	}, nil)
	if err != nil {
		return err
	}

	if attr, exists := resources.Attributes["description"]; !exists || attr.Expr == nil {
		// Use the first block's DefRange (resource definition range)
		var issueRange hcl.Range
		if len(resources.Blocks) > 0 {
			issueRange = resources.Blocks[0].DefRange
		} else if attr != nil {
			issueRange = attr.Range
		}
		
		runner.EmitIssue(
			r,
			"Workspace should have a description for governance and documentation",
			issueRange,
		)
	}

	return nil
}
