package apispec

import (
	"fmt"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// FabricEventhouseInvalidDescription checks whether fabric_eventhouse.description is valid
type FabricEventhouseInvalidDescription struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string

	maxLength int
}

// NewFabricRule returns a new rule instance
func NewFabricEventhouseInvalidDescription() *FabricEventhouseInvalidDescription {
	return &FabricEventhouseInvalidDescription{
		resourceType:  "fabric_eventhouse",
		attributeName: "description",

		maxLength: 1024,
	}
}

// Name returns the rule name
func (r *FabricEventhouseInvalidDescription) Name() string {
	return "fabric_eventhouse_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricEventhouseInvalidDescription) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricEventhouseInvalidDescription) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricEventhouseInvalidDescription) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricEventhouseInvalidDescription) Check(runner tflint.Runner) error {
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
				fmt.Sprintf("description must be at most %d characters (actual: %d)", r.maxLength, len(val)),
				attribute.Expr.Range(),
			)
		}

	}

	return nil
}
