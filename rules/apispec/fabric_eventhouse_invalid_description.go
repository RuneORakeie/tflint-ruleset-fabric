package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricEventhouseInvalidDescription struct{ tflint.DefaultRule }

func NewFabricEventhouseInvalidDescription() *FabricEventhouseInvalidDescription {
	return &FabricEventhouseInvalidDescription{}
}

func (r *FabricEventhouseInvalidDescription) Name() string {
	return "fabric_eventhouse_invalid_description"
}
func (r *FabricEventhouseInvalidDescription) Enabled() bool             { return true }
func (r *FabricEventhouseInvalidDescription) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricEventhouseInvalidDescription) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/eventhouse/definitions.json"
}

func (r *FabricEventhouseInvalidDescription) Check(runner tflint.Runner) error {
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
		if block.Labels[0] != "fabric_eventhouse" {
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
