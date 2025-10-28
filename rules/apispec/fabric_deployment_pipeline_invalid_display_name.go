package apispec

import (
	"fmt"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// FabricDeploymentPipelineInvalidDisplayName checks whether fabric_deployment_pipeline.display_name is valid
type FabricDeploymentPipelineInvalidDisplayName struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string

	maxLength int
}

// NewFabricRule returns a new rule instance
func NewFabricDeploymentPipelineInvalidDisplayName() *FabricDeploymentPipelineInvalidDisplayName {
	return &FabricDeploymentPipelineInvalidDisplayName{
		resourceType:  "fabric_deployment_pipeline",
		attributeName: "display_name",

		maxLength: 256,
	}
}

// Name returns the rule name
func (r *FabricDeploymentPipelineInvalidDisplayName) Name() string {
	return "fabric_deployment_pipeline_invalid_display_name"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricDeploymentPipelineInvalidDisplayName) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricDeploymentPipelineInvalidDisplayName) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricDeploymentPipelineInvalidDisplayName) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricDeploymentPipelineInvalidDisplayName) Check(runner tflint.Runner) error {
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
