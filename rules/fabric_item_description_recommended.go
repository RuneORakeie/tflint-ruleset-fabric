package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricItemDescriptionRecommended warns when items don't have descriptions
// Descriptions help document your Fabric environment, making it easier for teams to understand
// the purpose and ownership of resources, especially in large collaborative environments.
type FabricItemDescriptionRecommended struct {
	tflint.DefaultRule
}

func NewFabricItemDescriptionRecommended() *FabricItemDescriptionRecommended {
	return &FabricItemDescriptionRecommended{}
}

func (r *FabricItemDescriptionRecommended) Name() string {
	return "fabric_item_description_recommended"
}

func (r *FabricItemDescriptionRecommended) Enabled() bool {
	return true
}

func (r *FabricItemDescriptionRecommended) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *FabricItemDescriptionRecommended) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *FabricItemDescriptionRecommended) Check(runner tflint.Runner) error {
	// List of resources that have description attribute
	resourceTypes := []string{
		"fabric_activator",
		"fabric_copy_job",
		"fabric_data_pipeline",
		"fabric_dataflow",
		"fabric_deployment_pipeline",
		"fabric_digital_twin_builder",
		"fabric_domain",
		"fabric_eventhouse",
		"fabric_graphql_api",
		"fabric_kql_dashboard",
		"fabric_kql_database",
		"fabric_kql_queryset",
		"fabric_ml_experiment",
		"fabric_ml_model",
		"fabric_notebook",
		"fabric_spark_job_definition",
		"fabric_sql_database",
		"fabric_warehouse",
		"fabric_workspace",
	}

	for _, resourceType := range resourceTypes {
		resourceContent, err := runner.GetResourceContent(resourceType, &hclext.BodySchema{
			Attributes: []hclext.AttributeSchema{
				{Name: "description"},
			},
		}, nil)
		if err != nil {
			return err
		}

		for _, resource := range resourceContent.Blocks {
			if attr, exists := resource.Body.Attributes["description"]; !exists || attr.Expr == nil {
				runner.EmitIssue(
					r,
					"Adding a description improves documentation and governance of your Fabric environment. Consider including the purpose, owner, and any relevant business context.",
					resource.DefRange,
				)
			} else {
				// Check if description is empty string
				var description string
				if err := runner.EvaluateExpr(attr.Expr, &description, nil); err == nil && description == "" {
					runner.EmitIssue(
						r,
						"Description is empty. Consider adding meaningful information about the purpose, owner, and business context of this resource.",
						attr.Range,
					)
				}
			}
		}
	}

	return nil
}
