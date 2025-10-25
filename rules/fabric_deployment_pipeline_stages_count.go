package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricDeploymentPipelineStagesCount checks that deployment pipelines have between 2 and 10 stages
type FabricDeploymentPipelineStagesCount struct {
	tflint.DefaultRule
}

func NewFabricDeploymentPipelineStagesCount() *FabricDeploymentPipelineStagesCount {
	return &FabricDeploymentPipelineStagesCount{}
}

func (r *FabricDeploymentPipelineStagesCount) Name() string {
	return "fabric_deployment_pipeline_stages_count"
}

func (r *FabricDeploymentPipelineStagesCount) Enabled() bool {
	return true
}

func (r *FabricDeploymentPipelineStagesCount) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *FabricDeploymentPipelineStagesCount) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *FabricDeploymentPipelineStagesCount) Check(runner tflint.Runner) error {
	resourceContent, err := runner.GetResourceContent("fabric_deployment_pipeline", &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "stages",
				Body: &hclext.BodySchema{},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	const minStages = 2
	const maxStages = 10

	for _, resource := range resourceContent.Blocks {
		stagesBlocks := resource.Body.Blocks.OfType("stages")
		stageCount := len(stagesBlocks)

		if stageCount < minStages {
			runner.EmitIssue(
				r,
				fmt.Sprintf("Deployment pipeline must have at least %d stages (current: %d)", minStages, stageCount),
				resource.DefRange,
			)
		} else if stageCount > maxStages {
			runner.EmitIssue(
				r,
				fmt.Sprintf("Deployment pipeline must not exceed %d stages (current: %d)", maxStages, stageCount),
				resource.DefRange,
			)
		}
	}

	return nil
}
