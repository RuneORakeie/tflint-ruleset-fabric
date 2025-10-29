package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricApacheAirflowJobInvalidDescription struct{ tflint.DefaultRule }

func NewFabricApacheAirflowJobInvalidDescription() *FabricApacheAirflowJobInvalidDescription {
	return &FabricApacheAirflowJobInvalidDescription{}
}

func (r *FabricApacheAirflowJobInvalidDescription) Name() string {
	return "fabric_apache_airflow_job_invalid_description"
}
func (r *FabricApacheAirflowJobInvalidDescription) Enabled() bool             { return true }
func (r *FabricApacheAirflowJobInvalidDescription) Severity() tflint.Severity { return tflint.ERROR }
func (r *FabricApacheAirflowJobInvalidDescription) Link() string {
	return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/apacheAirflowJob/definitions.json"
}

func (r *FabricApacheAirflowJobInvalidDescription) Check(runner tflint.Runner) error {
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
		if block.Labels[0] != "fabric_apache_airflow_job" {
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
