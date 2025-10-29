package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricDigitalTwinBuilderInvalidDescription struct{ tflint.DefaultRule }

func NewFabricDigitalTwinBuilderInvalidDescription() *FabricDigitalTwinBuilderInvalidDescription {
	return &FabricDigitalTwinBuilderInvalidDescription{}
}

func (r *FabricDigitalTwinBuilderInvalidDescription) Name() string {
	return "fabric_digital_twin_builder_invalid_description"
}
func (r *FabricDigitalTwinBuilderInvalidDescription) Enabled() bool             { return true }
func (r *FabricDigitalTwinBuilderInvalidDescription) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricDigitalTwinBuilderInvalidDescription) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/digitalTwinBuilder/definitions.json"
}

func (r *FabricDigitalTwinBuilderInvalidDescription) Check(runner tflint.Runner) error {
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
		if block.Labels[0] != "fabric_digital_twin_builder" {
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
		if len(v) > 256 {
			if err := runner.EmitIssue(r,
				fmt.Sprintf("%s exceeds max length %d", "description", 256),
				attr.Expr.Range()); err != nil {
				return err
			}
		}
		// TODO: add pattern/enum checks if needed
	}

	return nil
}
