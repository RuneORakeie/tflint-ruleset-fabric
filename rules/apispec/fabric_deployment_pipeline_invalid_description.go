package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricDeploymentPipelineInvalidDescription struct{ tflint.DefaultRule }

func NewFabricDeploymentPipelineInvalidDescription() *FabricDeploymentPipelineInvalidDescription {
	return &FabricDeploymentPipelineInvalidDescription{}
}

func (r *FabricDeploymentPipelineInvalidDescription) Name() string {
	return "fabric_deployment_pipeline_invalid_description"
}
func (r *FabricDeploymentPipelineInvalidDescription) Enabled() bool             { return true }
func (r *FabricDeploymentPipelineInvalidDescription) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricDeploymentPipelineInvalidDescription) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/platform/definitions/deploymentPipelines.json"
}

func (r *FabricDeploymentPipelineInvalidDescription) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "resource",
				LabelNames: []string{"type", "name"},
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "description"},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, block := range content.Blocks {
		if block.Labels[0] != "fabric_deployment_pipeline" {
			continue
		}
		attr, ok := block.Body.Attributes["description"]
		if !ok {
			continue
		}

		var v string
		if err := runner.EvaluateExpr(attr.Expr, &v, nil); err != nil {
			continue
		}
		if len(v) > 1024 {
			if err := runner.EmitIssue(r,
				fmt.Sprintf("%s exceeds max length %d", "description", 1024),
				attr.Expr.Range()); err != nil {
				return err
			}
		}
		// TODO: add pattern/enum checks if needed
	}

	return nil
}
