package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricNotebookInvalidDisplayName checks whether fabric_notebook.display_name is valid
type FabricNotebookInvalidDisplayName struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string





	maxLength     int



}

// NewFabricRule returns a new rule instance
func NewFabricNotebookInvalidDisplayName() *FabricNotebookInvalidDisplayName {
	return &FabricNotebookInvalidDisplayName{
		resourceType:  "fabric_notebook",
		attributeName: "display_name",





		maxLength:     256,



	}
}

// Name returns the rule name
func (r *FabricNotebookInvalidDisplayName) Name() string {
	return "fabric_notebook_invalid_display_name"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricNotebookInvalidDisplayName) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricNotebookInvalidDisplayName) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricNotebookInvalidDisplayName) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricNotebookInvalidDisplayName) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}


		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)
		if err != nil {
			return err
		}




		if len(val) > r.maxLength {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("display_name must be at most %d characters (actual: %d)", r.maxLength, len(val)),
				attribute.Expr.Range(),
			)
		}



	}

	return nil
}




