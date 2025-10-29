package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricKQLDatabaseInvalidDescription struct{ tflint.DefaultRule }

func NewFabricKQLDatabaseInvalidDescription() *FabricKQLDatabaseInvalidDescription {
	return &FabricKQLDatabaseInvalidDescription{}
}

func (r *FabricKQLDatabaseInvalidDescription) Name() string {
	return "fabric_kql_database_invalid_description"
}
func (r *FabricKQLDatabaseInvalidDescription) Enabled() bool             { return true }
func (r *FabricKQLDatabaseInvalidDescription) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricKQLDatabaseInvalidDescription) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/kqlDatabase/definitions.json"
}

func (r *FabricKQLDatabaseInvalidDescription) Check(runner tflint.Runner) error {
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
		if block.Labels[0] != "fabric_kql_database" {
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
