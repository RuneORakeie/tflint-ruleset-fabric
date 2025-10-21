package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// FabricWorkspaceCapacity ensures capacity is assigned to workspaces
type FabricWorkspaceCapacity struct {
	tflint.DefaultRule
}

func NewFabricWorkspaceCapacity() *FabricWorkspaceCapacity {
	return &FabricWorkspaceCapacity{}
}

func (r *FabricWorkspaceCapacity) Name() string {
	return "fabric_workspace_capacity_required"
}

func (r *FabricWorkspaceCapacity) Enabled() bool {
	return true
}

func (r *FabricWorkspaceCapacity) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *FabricWorkspaceCapacity) Link() string {
	return "https://learn.microsoft.com/en-us/fabric/admin/capacity-settings"
}

func (r *FabricWorkspaceCapacity) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("fabric_workspace", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "capacity_id"},
		},
	}, nil)
	if err != nil {
		return err
	}

	if attr, exists := resources.Attributes["capacity_id"]; !exists || attr.Expr == nil {
		// Use the first block's DefRange (resource definition range)
		var issueRange hcl.Range
		if len(resources.Blocks) > 0 {
			issueRange = resources.Blocks[0].DefRange
		} else if attr != nil {
			issueRange = attr.Range
		}
		
		runner.EmitIssue(
			r,
			"Workspace should have a capacity assigned for production use",
			issueRange,
		)
	}

	return nil
}
