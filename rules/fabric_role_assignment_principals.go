package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// FabricRoleAssignmentPrincipals ensures role assignments specify principals
type FabricRoleAssignmentPrincipals struct {
	tflint.DefaultRule
}

func NewFabricRoleAssignmentPrincipals() *FabricRoleAssignmentPrincipals {
	return &FabricRoleAssignmentPrincipals{}
}

func (r *FabricRoleAssignmentPrincipals) Name() string {
	return "fabric_role_assignment_principal_required"
}

func (r *FabricRoleAssignmentPrincipals) Enabled() bool {
	return true
}

func (r *FabricRoleAssignmentPrincipals) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *FabricRoleAssignmentPrincipals) Link() string {
	return "https://learn.microsoft.com/en-us/fabric/admin/manage-user-permissions"
}

func (r *FabricRoleAssignmentPrincipals) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("fabric_workspace_role_assignment", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "principal_id"},
		},
	}, nil)
	if err != nil {
		return err
	}

	if attr, exists := resources.Attributes["principal_id"]; !exists || attr.Expr == nil {
		// Use the first block's DefRange (resource definition range)
		var issueRange hcl.Range
		if len(resources.Blocks) > 0 {
			issueRange = resources.Blocks[0].DefRange
		} else if attr != nil {
			issueRange = attr.Range
		}
		
		runner.EmitIssue(
			r,
			"Role assignment must specify a principal_id",
			issueRange,
		)
	}

	return nil
}
