package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricGraphqlAPIInvalidDescription struct{ tflint.DefaultRule }

func NewFabricGraphqlAPIInvalidDescription() *FabricGraphqlAPIInvalidDescription {
	return &FabricGraphqlAPIInvalidDescription{}
}

func (r *FabricGraphqlAPIInvalidDescription) Name() string {
	return "fabric_graphql_api_invalid_description"
}
func (r *FabricGraphqlAPIInvalidDescription) Enabled() bool             { return true }
func (r *FabricGraphqlAPIInvalidDescription) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricGraphqlAPIInvalidDescription) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/graphQLApi/definitions.json"
}

func (r *FabricGraphqlAPIInvalidDescription) Check(runner tflint.Runner) error {
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
		if block.Labels[0] != "fabric_graphql_api" {
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
