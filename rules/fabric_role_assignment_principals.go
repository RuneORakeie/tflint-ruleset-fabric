package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricRoleAssignmentRecommended warns when resources are created without role assignments
// Without role assignments, resources will only be accessible to admins or may not be accessible at all
// Applies to: workspaces, deployment pipelines, domains, and gateways
type FabricRoleAssignmentRecommended struct {
	tflint.DefaultRule
}

func NewFabricRoleAssignmentRecommended() *FabricRoleAssignmentRecommended {
	return &FabricRoleAssignmentRecommended{}
}

func (r *FabricRoleAssignmentRecommended) Name() string {
	return "fabric_role_assignment_recommended"
}

func (r *FabricRoleAssignmentRecommended) Enabled() bool {
	return true
}

func (r *FabricRoleAssignmentRecommended) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *FabricRoleAssignmentRecommended) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *FabricRoleAssignmentRecommended) Check(runner tflint.Runner) error {
	// Define resource types and their corresponding role assignment resources
	resourceConfigs := []struct {
		resourceType           string
		roleAssignmentType     string
		referenceAttribute     string
		displayNameAttribute   string
		resourceTypeFriendly   string
	}{
		{
			resourceType:         "fabric_workspace",
			roleAssignmentType:   "fabric_workspace_role_assignment",
			referenceAttribute:   "workspace_id",
			displayNameAttribute: "display_name",
			resourceTypeFriendly: "Workspace",
		},
		{
			resourceType:         "fabric_deployment_pipeline",
			roleAssignmentType:   "fabric_deployment_pipeline_role_assignment",
			referenceAttribute:   "deployment_pipeline_id",
			displayNameAttribute: "display_name",
			resourceTypeFriendly: "Deployment pipeline",
		},
		{
			resourceType:         "fabric_domain",
			roleAssignmentType:   "fabric_domain_role_assignment",
			referenceAttribute:   "domain_id",
			displayNameAttribute: "display_name",
			resourceTypeFriendly: "Domain",
		},
		{
			resourceType:         "fabric_gateway",
			roleAssignmentType:   "fabric_gateway_role_assignment",
			referenceAttribute:   "gateway_id",
			displayNameAttribute: "display_name",
			resourceTypeFriendly: "Gateway",
		},
	}

	for _, config := range resourceConfigs {
		if err := r.checkResourceRoleAssignments(runner, config); err != nil {
			return err
		}
	}

	return nil
}

type resourceConfig struct {
	resourceType           string
	roleAssignmentType     string
	referenceAttribute     string
	displayNameAttribute   string
	resourceTypeFriendly   string
}

func (r *FabricRoleAssignmentRecommended) checkResourceRoleAssignments(
	runner tflint.Runner,
	config struct {
		resourceType           string
		roleAssignmentType     string
		referenceAttribute     string
		displayNameAttribute   string
		resourceTypeFriendly   string
	},
) error {
	// Get all resources of this type
	resourceSchema := &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: config.displayNameAttribute},
		},
	}
	
	resources, err := runner.GetResourceContent(config.resourceType, resourceSchema, nil)
	if err != nil {
		return err
	}

	// Get all role assignment resources for this type
	roleAssignmentSchema := &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: config.referenceAttribute},
		},
	}
	
	roleAssignments, err := runner.GetResourceContent(config.roleAssignmentType, roleAssignmentSchema, nil)
	if err != nil {
		return err
	}

	// Build a set of resource IDs that have role assignments
	resourcesWithRoles := make(map[string]bool)
	
	for _, block := range roleAssignments.Blocks {
		if attr, exists := block.Body.Attributes[config.referenceAttribute]; exists {
			var resourceRef string
			// Try to evaluate the expression to see if it references a resource
			err := runner.EvaluateExpr(attr.Expr, &resourceRef, nil)
			if err == nil {
				resourcesWithRoles[resourceRef] = true
			}
		}
	}

	// Check each resource
	for _, block := range resources.Blocks {
		// Get the resource reference (e.g., "fabric_workspace.example")
		resourceRef := fmt.Sprintf("%s.%s", config.resourceType, block.Labels[0])
		
		// If this resource doesn't have any role assignments, emit a warning
		if !resourcesWithRoles[resourceRef] {
			var displayName string
			if attr, exists := block.Body.Attributes[config.displayNameAttribute]; exists {
				_ = runner.EvaluateExpr(attr.Expr, &displayName, nil)
			}
			
			message := fmt.Sprintf("%s '%s' does not have any role assignments. This resource may not be accessible to users.", 
				config.resourceTypeFriendly, block.Labels[0])
			if displayName != "" {
				message = fmt.Sprintf("%s '%s' (%s) does not have any role assignments. This resource may not be accessible to users.", 
					config.resourceTypeFriendly, displayName, block.Labels[0])
			}
			
			runner.EmitIssue(
				r,
				message,
				block.DefRange,
			)
		}
	}

	return nil
}
