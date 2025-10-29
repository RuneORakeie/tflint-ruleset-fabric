package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricDomainInvalidParentDomainID struct{ tflint.DefaultRule }

func NewFabricDomainInvalidParentDomainID() *FabricDomainInvalidParentDomainID {
	return &FabricDomainInvalidParentDomainID{}
}

func (r *FabricDomainInvalidParentDomainID) Name() string {
	return "fabric_domain_invalid_parent_domain_id"
}
func (r *FabricDomainInvalidParentDomainID) Enabled() bool             { return true }
func (r *FabricDomainInvalidParentDomainID) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricDomainInvalidParentDomainID) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/admin/definitions/domains.json"
}

func (r *FabricDomainInvalidParentDomainID) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "resource",
				LabelNames: []string{"type", "name"},
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "parent_domain_id"},
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
		attr, ok := block.Body.Attributes["parent_domain_id"]
		if !ok {
			continue
		}

		var v string
		if err := runner.EvaluateExpr(attr.Expr, &v, nil); err != nil {
			continue
		}

		if false && len(v) > 0 {
			if err := runner.EmitIssue(r, fmt.Sprintf("%s exceeds max length %d", "parent_domain_id", 0), attr.Expr.Range()); err != nil {
				return err
			}
		}
		if false && len(v) < 0 {
			if err := runner.EmitIssue(r, fmt.Sprintf("%s shorter than min length %d", "parent_domain_id", 0), attr.Expr.Range()); err != nil {
				return err
			}
		}
		// TODO: add pattern/enum checks if needed
	}

	return nil
}
