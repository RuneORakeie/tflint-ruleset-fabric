package apispec

import (
	"fmt"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// FabricDeploymentPipelineInvalidDescription checks whether fabric_deployment_pipeline.description is valid
type FabricDeploymentPipelineInvalidDescription struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string

	maxLength int
}

// NewFabricRule returns a new rule instance
func NewFabricDeploymentPipelineInvalidDescription() *FabricDeploymentPipelineInvalidDescription {
	return &FabricDeploymentPipelineInvalidDescription{
		resourceType:  "fabric_deployment_pipeline",
		attributeName: "description",

		maxLength: 1024,
	}
}

// Name returns the rule name
func (r *FabricDeploymentPipelineInvalidDescription) Name() string {
	return "fabric_deployment_pipeline_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricDeploymentPipelineInvalidDescription) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricDeploymentPipelineInvalidDescription) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricDeploymentPipelineInvalidDescription) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricDeploymentPipelineInvalidDescription) Check(runner tflint.Runner) error {
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
