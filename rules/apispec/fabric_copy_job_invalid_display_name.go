package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricCopyJobInvalidDisplayName struct{ tflint.DefaultRule }

func NewFabricCopyJobInvalidDisplayName() *FabricCopyJobInvalidDisplayName {
	return &FabricCopyJobInvalidDisplayName{}
}

func (r *FabricCopyJobInvalidDisplayName) Name() string {
	return "fabric_copy_job_invalid_display_name"
}
func (r *FabricCopyJobInvalidDisplayName) Enabled() bool             { return true }
func (r *FabricCopyJobInvalidDisplayName) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricCopyJobInvalidDisplayName) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/copyJob/definitions.json"
}

func (r *FabricCopyJobInvalidDisplayName) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "resource",
				LabelNames: []string{"type", "name"},
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

	for _, block := range content.Blocks {
		if block.Labels[0] != "fabric_copy_job" {
			continue
		}
		attr, ok := block.Body.Attributes["display_name"]
		if !ok {
			continue
		}

		var v string
		if err := runner.EvaluateExpr(attr.Expr, &v, nil); err != nil {
			continue
		}
		if len(v) > 256 {
			if err := runner.EmitIssue(r,
				fmt.Sprintf("%s exceeds max length %d", "display_name", 256),
				attr.Expr.Range()); err != nil {
				return err
			}
		}
		// TODO: add pattern/enum checks if needed
	}

	return nil
}
