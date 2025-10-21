package rules

import (
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// FabricWorkspaceNaming checks that workspace names follow naming conventions
type FabricWorkspaceNaming struct {
	tflint.DefaultRule
}

func NewFabricWorkspaceNaming() *FabricWorkspaceNaming {
	return &FabricWorkspaceNaming{}
}

func (r *FabricWorkspaceNaming) Name() string {
	return "fabric_workspace_naming"
}

func (r *FabricWorkspaceNaming) Enabled() bool {
	return true
}

func (r *FabricWorkspaceNaming) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *FabricWorkspaceNaming) Link() string {
	return "https://learn.microsoft.com/en-us/fabric/admin/fabric-governance"
}

func (r *FabricWorkspaceNaming) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("fabric_workspace", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "display_name"},
		},
	}, nil)
	if err != nil {
		return err
	}

	namePattern := regexp.MustCompile(`^[a-z0-9\-]{3,50}$`)

	if attr, exists := resources.Attributes["display_name"]; exists && attr.Expr != nil {
		var name string
		if err := runner.EvaluateExpr(attr.Expr, &name, nil); err == nil && name != "" {
			if !namePattern.MatchString(name) {
				runner.EmitIssue(
					r,
					"Workspace name should be 3-50 characters and contain only lowercase letters, numbers, and hyphens",
					attr.Range,
				)
			}
		}
	}

	return nil
}
