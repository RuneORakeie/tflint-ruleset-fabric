package apispec

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricGatewayInvalidType checks whether fabric_gateway.type is valid
type FabricGatewayInvalidType struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string








}

// NewFabricRule returns a new rule instance
func NewFabricGatewayInvalidType() *FabricGatewayInvalidType {
	return &FabricGatewayInvalidType{
		resourceType:  "fabric_gateway",
		attributeName: "type",
		enum:          []string{ "VirtualNetwork",  },








	}
}

// Name returns the rule name
func (r *FabricGatewayInvalidType) Name() string {
	return "fabric_gateway_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricGatewayInvalidType) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricGatewayInvalidType) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricGatewayInvalidType) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricGatewayInvalidType) Check(runner tflint.Runner) error {
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

func (r *FabricGatewayInvalidType) validateEnum(runner tflint.Runner, val string, attribute *hclext.Attribute) error {
	for _, valid := range r.enum {
		if val == valid {
			return nil
		}
	}
	return runner.EmitIssue(
		r,
		fmt.Sprintf("type must be one of: %v", r.enum),
		attribute.Expr.Range(),
	)
}




