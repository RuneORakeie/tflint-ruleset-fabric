package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hcl"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
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

func (r *FabricRoleAssignmentPrincipals) Severity() string {
	return tflint.ERROR
}

func (r *FabricRoleAssignmentPrincipals) Link() string {
	return "https://learn.microsoft.com/en-us/fabric/admin/manage-user-permissions"
}

func (r *FabricRoleAssignmentPrincipals) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourcesByType("fabric_workspace_role_assignment")
	if err != nil {
		return err
	}

	for _, resource := range resources {
		var principalID string
		var principal hcl.Attribute

		err := resource.GetAttribute("principal_id", &principalID)
		if err != nil {
			// Try getting as block
			principalBlock := resource.GetBlock("principal")
			if principalBlock == nil || len(principalBlock) == 0 {
				runner.EmitIssue(
					r,
					"Role assignment must specify a principal_id or principal block",
					resource.GetNameRange(),
				)
			}
		}
	}

	return nil
}