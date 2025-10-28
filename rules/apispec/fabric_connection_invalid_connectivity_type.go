package apispec

import (
	"fmt"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// FabricConnectionInvalidConnectivityType checks whether fabric_connection.connectivity_type is valid
type FabricConnectionInvalidConnectivityType struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewFabricRule returns a new rule instance
func NewFabricConnectionInvalidConnectivityType() *FabricConnectionInvalidConnectivityType {
	return &FabricConnectionInvalidConnectivityType{
		resourceType:  "fabric_connection",
		attributeName: "connectivity_type",
		enum:          []string{"ShareableCloud", "VirtualNetworkGateway"},
	}
}

// Name returns the rule name
func (r *FabricConnectionInvalidConnectivityType) Name() string {
	return "fabric_connection_invalid_connectivity_type"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricConnectionInvalidConnectivityType) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricConnectionInvalidConnectivityType) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricConnectionInvalidConnectivityType) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricConnectionInvalidConnectivityType) Check(runner tflint.Runner) error {
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

func (r *FabricConnectionInvalidConnectivityType) validateEnum(runner tflint.Runner, val string, attribute *hclext.Attribute) error {
	for _, valid := range r.enum {
		if val == valid {
			return nil
		}
	}
	return runner.EmitIssue(
		r,
		fmt.Sprintf("connectivity_type must be one of: %v", r.enum),
		attribute.Expr.Range(),
	)
}
