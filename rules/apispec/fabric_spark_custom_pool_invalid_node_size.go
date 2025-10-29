package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricSparkCustomPoolInvalidNodeSize struct{ tflint.DefaultRule }

func NewFabricSparkCustomPoolInvalidNodeSize() *FabricSparkCustomPoolInvalidNodeSize {
	return &FabricSparkCustomPoolInvalidNodeSize{}
}

func (r *FabricSparkCustomPoolInvalidNodeSize) Name() string {
	return "fabric_spark_custom_pool_invalid_node_size"
}
func (r *FabricSparkCustomPoolInvalidNodeSize) Enabled() bool             { return true }
func (r *FabricSparkCustomPoolInvalidNodeSize) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricSparkCustomPoolInvalidNodeSize) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/spark/definitions.json"
}

func (r *FabricSparkCustomPoolInvalidNodeSize) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "resource",
				LabelNames: []string{"type", "name"},
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "node_size"},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, block := range content.Blocks {
		if block.Labels[0] != "fabric_spark_custom_pool" {
			continue
		}
		attr, ok := block.Body.Attributes["node_size"]
		if !ok {
			continue
		}

		var v string
		if err := runner.EvaluateExpr(attr.Expr, &v, nil); err != nil {
			continue
		}

		if false && len(v) > 0 {
			if err := runner.EmitIssue(r, fmt.Sprintf("%s exceeds max length %d", "node_size", 0), attr.Expr.Range()); err != nil {
				return err
			}
		}
		if false && len(v) < 0 {
			if err := runner.EmitIssue(r, fmt.Sprintf("%s shorter than min length %d", "node_size", 0), attr.Expr.Range()); err != nil {
				return err
			}
		}
		// TODO: add pattern/enum checks if needed
	}

	return nil
}
