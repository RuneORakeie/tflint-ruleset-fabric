package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricDomainInvalidDescription checks whether fabric_domain.description is valid
type FabricDomainInvalidDescription struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string





	maxLength     int



}

// NewFabricRule returns a new rule instance
func NewFabricDomainInvalidDescription() *FabricDomainInvalidDescription {
	return &FabricDomainInvalidDescription{
		resourceType:  "fabric_domain",
		attributeName: "description",





		maxLength:     256,



	}
}

// Name returns the rule name
func (r *FabricDomainInvalidDescription) Name() string {
	return "fabric_domain_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricDomainInvalidDescription) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricDomainInvalidDescription) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricDomainInvalidDescription) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricDomainInvalidDescription) Check(runner tflint.Runner) error {
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




