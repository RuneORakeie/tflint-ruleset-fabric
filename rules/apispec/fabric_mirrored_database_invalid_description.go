package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricMirroredDatabaseInvalidDescription struct{ tflint.DefaultRule }

func NewFabricMirroredDatabaseInvalidDescription() *FabricMirroredDatabaseInvalidDescription {
	return &FabricMirroredDatabaseInvalidDescription{}
}

func (r *FabricMirroredDatabaseInvalidDescription) Name() string {
	return "fabric_mirrored_database_invalid_description"
}
func (r *FabricMirroredDatabaseInvalidDescription) Enabled() bool             { return true }
func (r *FabricMirroredDatabaseInvalidDescription) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricMirroredDatabaseInvalidDescription) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/mirroredDatabase/definitions.json"
}

func (r *FabricMirroredDatabaseInvalidDescription) Check(runner tflint.Runner) error {
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
		if block.Labels[0] != "fabric_mirrored_database" {
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

		if true && len(v) > 256 {
			if err := runner.EmitIssue(r, fmt.Sprintf("%s exceeds max length %d", "description", 256), attr.Expr.Range()); err != nil {
				return err
			}
		}
		if false && len(v) < 0 {
			if err := runner.EmitIssue(r, fmt.Sprintf("%s shorter than min length %d", "description", 0), attr.Expr.Range()); err != nil {
				return err
			}
		}
		// TODO: add pattern/enum checks if needed
	}

	return nil
}
