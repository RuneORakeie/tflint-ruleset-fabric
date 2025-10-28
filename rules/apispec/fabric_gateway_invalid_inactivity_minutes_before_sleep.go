package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)


// FabricGatewayInvalidInactivityMinutesBeforeSleep checks whether fabric_gateway.inactivity_minutes_before_sleep is valid
type FabricGatewayInvalidInactivityMinutesBeforeSleep struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewFabricRule returns a new rule instance
func NewFabricGatewayInvalidInactivityMinutesBeforeSleep() *FabricGatewayInvalidInactivityMinutesBeforeSleep {
	return &FabricGatewayInvalidInactivityMinutesBeforeSleep{
		resourceType:  "fabric_gateway",
		attributeName: "inactivity_minutes_before_sleep",
		enum:          []string{"30", "60", "90", "120", "150", "240", "360", "480", "720", "1440"},
	}
}

// Name returns the rule name
func (r *FabricGatewayInvalidInactivityMinutesBeforeSleep) Name() string {
	return "fabric_gateway_invalid_inactivity_minutes_before_sleep"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricGatewayInvalidInactivityMinutesBeforeSleep) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricGatewayInvalidInactivityMinutesBeforeSleep) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricGatewayInvalidInactivityMinutesBeforeSleep) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricGatewayInvalidInactivityMinutesBeforeSleep) Check(runner tflint.Runner) error {
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

		if err := r.validateEnum(runner, val, attribute); err != nil {
			return err
		}

	}

	return nil
}

func (r *FabricGatewayInvalidInactivityMinutesBeforeSleep) validateEnum(runner tflint.Runner, val string, attribute *hclext.Attribute) error {
	for _, valid := range r.enum {
		if val == valid {
			return nil
		}
	}
	return runner.EmitIssue(
		r,
		fmt.Sprintf("inactivity_minutes_before_sleep must be one of: %v", r.enum),
		attribute.Expr.Range(),
	)
}
