package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricCopyJobInvalidDisplayName checks whether fabric_copy_job.display_name is valid
type FabricCopyJobInvalidDisplayName struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string

	maxLength int
}

// NewFabricRule returns a new rule instance
func NewFabricCopyJobInvalidDisplayName() *FabricCopyJobInvalidDisplayName {
	return &FabricCopyJobInvalidDisplayName{
		resourceType:  "fabric_copy_job",
		attributeName: "display_name",

		maxLength: 256,
	}
}

// Name returns the rule name
func (r *FabricCopyJobInvalidDisplayName) Name() string {
	return "fabric_copy_job_invalid_display_name"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricCopyJobInvalidDisplayName) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricCopyJobInvalidDisplayName) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricCopyJobInvalidDisplayName) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricCopyJobInvalidDisplayName) Check(runner tflint.Runner) error {
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
