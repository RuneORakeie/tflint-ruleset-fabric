package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricConnectionInvalidPrivacyLevel struct{ tflint.DefaultRule }

func NewFabricConnectionInvalidPrivacyLevel() *FabricConnectionInvalidPrivacyLevel {
	return &FabricConnectionInvalidPrivacyLevel{}
}

func (r *FabricConnectionInvalidPrivacyLevel) Name() string {
	return "fabric_connection_invalid_privacy_level"
}
func (r *FabricConnectionInvalidPrivacyLevel) Enabled() bool             { return true }
func (r *FabricConnectionInvalidPrivacyLevel) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricConnectionInvalidPrivacyLevel) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/platform/definitions/connections.json"
}

func (r *FabricConnectionInvalidPrivacyLevel) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "resource",
				LabelNames: []string{"type", "name"},
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "privacy_level"},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, block := range content.Blocks {
		if block.Labels[0] != "fabric_connection" {
			continue
		}
		attr, ok := block.Body.Attributes["privacy_level"]
		if !ok {
			continue
		}

		var v string
		if err := runner.EvaluateExpr(attr.Expr, &v, nil); err != nil {
			continue
		}

		if false && len(v) > 0 {
			if err := runner.EmitIssue(r, fmt.Sprintf("%s exceeds max length %d", "privacy_level", 0), attr.Expr.Range()); err != nil {
				return err
			}
		}
		if false && len(v) < 0 {
			if err := runner.EmitIssue(r, fmt.Sprintf("%s shorter than min length %d", "privacy_level", 0), attr.Expr.Range()); err != nil {
				return err
			}
		}
		// TODO: add pattern/enum checks if needed
	}

	return nil
}
