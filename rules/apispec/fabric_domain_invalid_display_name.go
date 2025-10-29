package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricDomainInvalidDisplayName struct{ tflint.DefaultRule }

func NewFabricDomainInvalidDisplayName() *FabricDomainInvalidDisplayName {
	return &FabricDomainInvalidDisplayName{}
}

func (r *FabricDomainInvalidDisplayName) Name() string              { return "fabric_domain_invalid_display_name" }
func (r *FabricDomainInvalidDisplayName) Enabled() bool             { return true }
func (r *FabricDomainInvalidDisplayName) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricDomainInvalidDisplayName) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/admin/definitions/domains.json"
}

func (r *FabricDomainInvalidDisplayName) Check(runner tflint.Runner) error {
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
		if block.Labels[0] != "fabric_domain" {
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

		if true && len(v) > 40 {
			if err := runner.EmitIssue(r, fmt.Sprintf("%s exceeds max length %d", "display_name", 40), attr.Expr.Range()); err != nil {
				return err
			}
		}
		if false && len(v) < 0 {
			if err := runner.EmitIssue(r, fmt.Sprintf("%s shorter than min length %d", "display_name", 0), attr.Expr.Range()); err != nil {
				return err
			}
		}
		// TODO: add pattern/enum checks if needed
	}

	return nil
}
