package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricEventhouseInvalidFormat struct{ tflint.DefaultRule }

func NewFabricEventhouseInvalidFormat() *FabricEventhouseInvalidFormat {
	return &FabricEventhouseInvalidFormat{}
}

func (r *FabricEventhouseInvalidFormat) Name() string              { return "fabric_eventhouse_invalid_format" }
func (r *FabricEventhouseInvalidFormat) Enabled() bool             { return true }
func (r *FabricEventhouseInvalidFormat) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricEventhouseInvalidFormat) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/eventhouse/definitions.json"
}

func (r *FabricEventhouseInvalidFormat) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "resource",
				LabelNames: []string{"type", "name"},
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "format"},
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
		attr, ok := block.Body.Attributes["format"]
		if !ok {
			continue
		}

		var v string
		if err := runner.EvaluateExpr(attr.Expr, &v, nil); err != nil {
			continue
		}

		if false && len(v) > 0 {
			if err := runner.EmitIssue(r, fmt.Sprintf("%s exceeds max length %d", "format", 0), attr.Expr.Range()); err != nil {
				return err
			}
		}
		if false && len(v) < 0 {
			if err := runner.EmitIssue(r, fmt.Sprintf("%s shorter than min length %d", "format", 0), attr.Expr.Range()); err != nil {
				return err
			}
		}
		// TODO: add pattern/enum checks if needed
	}

	return nil
}
