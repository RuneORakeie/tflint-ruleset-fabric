package rules

import (
	"fmt"

	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// FabricDeploymentPipelineStagesDisplayNameLength checks stage display names don't exceed 256 chars
type FabricDeploymentPipelineStagesDisplayNameLength struct {
	tflint.DefaultRule
}

func NewFabricDeploymentPipelineStagesDisplayNameLength() *FabricDeploymentPipelineStagesDisplayNameLength {
	return &FabricDeploymentPipelineStagesDisplayNameLength{}
}

func (r *FabricDeploymentPipelineStagesDisplayNameLength) Name() string {
	return "fabric_deployment_pipeline_stages_display_name_length"
}

func (r *FabricDeploymentPipelineStagesDisplayNameLength) Enabled() bool {
	return true
}

func (r *FabricDeploymentPipelineStagesDisplayNameLength) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *FabricDeploymentPipelineStagesDisplayNameLength) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *FabricDeploymentPipelineStagesDisplayNameLength) Check(runner tflint.Runner) error {
	resourceContent, err := runner.GetResourceContent("fabric_deployment_pipeline", &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "stages",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "display_name"},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	const maxLength = 256

	for _, resource := range resourceContent.Blocks {
		stagesBlocks := resource.Body.Blocks.OfType("stages")

		for _, stage := range stagesBlocks {
			if attr, exists := stage.Body.Attributes["display_name"]; exists && attr.Expr != nil {
				var name string
				if err := runner.EvaluateExpr(attr.Expr, &name, nil); err == nil && name != "" {
					if len(name) > maxLength {
						runner.EmitIssue(
							r,
							fmt.Sprintf("Stage display_name must not exceed %d characters (current: %d)", maxLength, len(name)),
							attr.Range,
						)
					}
				}
			}
		}
	}

	return nil
}
