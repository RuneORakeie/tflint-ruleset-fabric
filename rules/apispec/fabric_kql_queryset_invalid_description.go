package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)


// FabricKQLQuerysetInvalidDescription checks whether fabric_kql_queryset.description is valid
type FabricKQLQuerysetInvalidDescription struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string

	maxLength int
}

// NewFabricRule returns a new rule instance
func NewFabricKQLQuerysetInvalidDescription() *FabricKQLQuerysetInvalidDescription {
	return &FabricKQLQuerysetInvalidDescription{
		resourceType:  "fabric_kql_queryset",
		attributeName: "description",

		maxLength: 256,
	}
}

// Name returns the rule name
func (r *FabricKQLQuerysetInvalidDescription) Name() string {
	return "fabric_kql_queryset_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricKQLQuerysetInvalidDescription) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricKQLQuerysetInvalidDescription) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricKQLQuerysetInvalidDescription) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricKQLQuerysetInvalidDescription) Check(runner tflint.Runner) error {
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
