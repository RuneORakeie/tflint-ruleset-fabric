package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricEventstreamInvalidDescription struct{ tflint.DefaultRule }

func NewFabricEventstreamInvalidDescription() *FabricEventstreamInvalidDescription {
	return &FabricEventstreamInvalidDescription{}
}

func (r *FabricEventstreamInvalidDescription) Name() string {
	return "fabric_eventstream_invalid_description"
}
func (r *FabricEventstreamInvalidDescription) Enabled() bool             { return true }
func (r *FabricEventstreamInvalidDescription) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricEventstreamInvalidDescription) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/eventstream/definitions.json"
}

func (r *FabricEventstreamInvalidDescription) Check(runner tflint.Runner) error {
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
		if block.Labels[0] != "fabric_eventstream" {
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
