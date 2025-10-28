package apispec

import (
	"fmt"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// FabricCopyJobInvalidDescription checks whether fabric_copy_job.description is valid
type FabricCopyJobInvalidDescription struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string

	maxLength int
}

// NewFabricRule returns a new rule instance
func NewFabricCopyJobInvalidDescription() *FabricCopyJobInvalidDescription {
	return &FabricCopyJobInvalidDescription{
		resourceType:  "fabric_copy_job",
		attributeName: "description",

		maxLength: 1021,
	}
}

// Name returns the rule name
func (r *FabricCopyJobInvalidDescription) Name() string {
	return "fabric_copy_job_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricCopyJobInvalidDescription) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricCopyJobInvalidDescription) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricCopyJobInvalidDescription) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricCopyJobInvalidDescription) Check(runner tflint.Runner) error {
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
