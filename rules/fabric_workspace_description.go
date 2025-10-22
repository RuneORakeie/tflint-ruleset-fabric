package rules

import (
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
	resourceContent, err := runner.GetResourceContent("fabric_workspace", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "description"},
		},
	}, nil)
	if err != nil {
		return err
	}

	// Iterate over all resource blocks
	for _, resource := range resourceContent.Blocks {
		if attr, exists := resource.Body.Attributes["description"]; !exists || attr.Expr == nil {
			runner.EmitIssue(
				r,
				"Workspace should have a description for governance and documentation",
				resource.DefRange,
			)
		}
	}

	return nil
}
