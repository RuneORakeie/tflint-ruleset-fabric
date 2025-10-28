package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)


// FabricDataPipelineInvalidDisplayName checks whether fabric_data_pipeline.display_name is valid
type FabricDataPipelineInvalidDisplayName struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string

	maxLength int
}

// NewFabricRule returns a new rule instance
func NewFabricDataPipelineInvalidDisplayName() *FabricDataPipelineInvalidDisplayName {
	return &FabricDataPipelineInvalidDisplayName{
		resourceType:  "fabric_data_pipeline",
		attributeName: "display_name",

		maxLength: 256,
	}
}

// Name returns the rule name
func (r *FabricDataPipelineInvalidDisplayName) Name() string {
	return "fabric_data_pipeline_invalid_display_name"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricDataPipelineInvalidDisplayName) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricDataPipelineInvalidDisplayName) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricDataPipelineInvalidDisplayName) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricDataPipelineInvalidDisplayName) Check(runner tflint.Runner) error {
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
