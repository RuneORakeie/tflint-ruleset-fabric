package rules

import (
	"fmt"
	"strings"

	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// FabricWorkspaceRoleAssignmentRole validates workspace role values
type FabricWorkspaceRoleAssignmentRole struct {
	tflint.DefaultRule
}

func NewFabricWorkspaceRoleAssignmentRole() *FabricWorkspaceRoleAssignmentRole {
	return &FabricWorkspaceRoleAssignmentRole{}
}

func (r *FabricWorkspaceRoleAssignmentRole) Name() string {
	return "fabric_workspace_role_assignment_role"
}

func (r *FabricWorkspaceRoleAssignmentRole) Enabled() bool {
	return true
}

func (r *FabricWorkspaceRoleAssignmentRole) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *FabricWorkspaceRoleAssignmentRole) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *FabricWorkspaceRoleAssignmentRole) Check(runner tflint.Runner) error {
	resourceContent, err := runner.GetResourceContent("fabric_workspace_role_assignment", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "role"},
		},
	}, nil)
	if err != nil {
		return err
	}

	validRoles := map[string]bool{
		"Admin":       true,
		"Contributor": true,
		"Member":      true,
		"Viewer":      true,
	}

	for _, resource := range resourceContent.Blocks {
		if attr, exists := resource.Body.Attributes["role"]; exists && attr.Expr != nil {
			var role string
			if err := runner.EvaluateExpr(attr.Expr, &role, nil); err == nil && role != "" {
				if !validRoles[role] {
					validRolesList := []string{"Admin", "Contributor", "Member", "Viewer"}
					runner.EmitIssue(
						r,
						fmt.Sprintf("Invalid workspace role '%s'. Must be one of: %s", role, strings.Join(validRolesList, ", ")),
						attr.Range,
					)
				}
			}
		}
	}

	return nil
}
