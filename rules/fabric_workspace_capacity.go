package rules

import (
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
	resourceContent, err := runner.GetResourceContent("fabric_workspace", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "capacity_id"},
		},
	}, nil)
	if err != nil {
		return err
	}

	// Iterate over all resource blocks
	for _, resource := range resourceContent.Blocks {
		if attr, exists := resource.Body.Attributes["capacity_id"]; !exists || attr.Expr == nil {
			runner.EmitIssue(
				r,
				"Workspace should have a capacity assigned for production use",
				resource.DefRange,
			)
		}
	}

	return nil
}
