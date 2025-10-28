package apispec

import (
	"fmt"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// FabricDomainInvalidDisplayName checks whether fabric_domain.display_name is valid
type FabricDomainInvalidDisplayName struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string

	maxLength int
}

// NewFabricRule returns a new rule instance
func NewFabricDomainInvalidDisplayName() *FabricDomainInvalidDisplayName {
	return &FabricDomainInvalidDisplayName{
		resourceType:  "fabric_domain",
		attributeName: "display_name",

		maxLength: 40,
	}
}

// Name returns the rule name
func (r *FabricDomainInvalidDisplayName) Name() string {
	return "fabric_domain_invalid_display_name"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricDomainInvalidDisplayName) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricDomainInvalidDisplayName) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricDomainInvalidDisplayName) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricDomainInvalidDisplayName) Check(runner tflint.Runner) error {
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
