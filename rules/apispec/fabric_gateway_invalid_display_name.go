package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)


// FabricGatewayInvalidDisplayName checks whether fabric_gateway.display_name is valid
type FabricGatewayInvalidDisplayName struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string

	maxLength int
}

// NewFabricRule returns a new rule instance
func NewFabricGatewayInvalidDisplayName() *FabricGatewayInvalidDisplayName {
	return &FabricGatewayInvalidDisplayName{
		resourceType:  "fabric_gateway",
		attributeName: "display_name",

		maxLength: 200,
	}
}

// Name returns the rule name
func (r *FabricGatewayInvalidDisplayName) Name() string {
	return "fabric_gateway_invalid_display_name"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricGatewayInvalidDisplayName) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricGatewayInvalidDisplayName) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricGatewayInvalidDisplayName) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricGatewayInvalidDisplayName) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)
		if err != nil {
			return err
		}

		if len(val) > r.maxLength {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("display_name must be at most %d characters (actual: %d)", r.maxLength, len(val)),
				attribute.Expr.Range(),
			)
		}

	}

	return nil
}
