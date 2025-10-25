package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricDeploymentPipelineStagesDescriptionLength checks stage descriptions don't exceed 1024 chars
type FabricDeploymentPipelineStagesDescriptionLength struct {
	tflint.DefaultRule
}

func NewFabricDeploymentPipelineStagesDescriptionLength() *FabricDeploymentPipelineStagesDescriptionLength {
	return &FabricDeploymentPipelineStagesDescriptionLength{}
}

func (r *FabricDeploymentPipelineStagesDescriptionLength) Name() string {
	return "fabric_deployment_pipeline_stages_description_length"
}

func (r *FabricDeploymentPipelineStagesDescriptionLength) Enabled() bool {
	return true
}

func (r *FabricDeploymentPipelineStagesDescriptionLength) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *FabricDeploymentPipelineStagesDescriptionLength) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *FabricDeploymentPipelineStagesDescriptionLength) Check(runner tflint.Runner) error {
	resourceContent, err := runner.GetResourceContent("fabric_deployment_pipeline", &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "stages",
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

	const maxLength = 1024

	for _, resource := range resourceContent.Blocks {
		stagesBlocks := resource.Body.Blocks.OfType("stages")
		
		for _, stage := range stagesBlocks {
			if attr, exists := stage.Body.Attributes["description"]; exists && attr.Expr != nil {
				var description string
				if err := runner.EvaluateExpr(attr.Expr, &description, nil); err == nil && description != "" {
					if len(description) > maxLength {
						runner.EmitIssue(
							r,
							fmt.Sprintf("Stage description must not exceed %d characters (current: %d)", maxLength, len(description)),
							attr.Range,
						)
					}
				}
			}
		}
	}

	return nil
}
