package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricCopyJobInvalidDescription struct{ tflint.DefaultRule }

func NewFabricCopyJobInvalidDescription() *FabricCopyJobInvalidDescription {
	return &FabricCopyJobInvalidDescription{}
}

func (r *FabricCopyJobInvalidDescription) Name() string              { return "fabric_copy_job_invalid_description" }
func (r *FabricCopyJobInvalidDescription) Enabled() bool             { return true }
func (r *FabricCopyJobInvalidDescription) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricCopyJobInvalidDescription) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/copyJob/definitions.json"
}

func (r *FabricCopyJobInvalidDescription) Check(runner tflint.Runner) error {
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
		if block.Labels[0] != "fabric_copy_job" {
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
		if len(v) > 1021 {
			if err := runner.EmitIssue(r,
				fmt.Sprintf("%s exceeds max length %d", "description", 1021),
				attr.Expr.Range()); err != nil {
				return err
			}
		}
		// TODO: add pattern/enum checks if needed
	}

	return nil
}
