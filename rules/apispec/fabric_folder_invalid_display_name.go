package apispec

import (
	"fmt"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// FabricFolderInvalidDisplayName checks whether fabric_folder.display_name is valid
type FabricFolderInvalidDisplayName struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string

	maxLength int
}

// NewFabricRule returns a new rule instance
func NewFabricFolderInvalidDisplayName() *FabricFolderInvalidDisplayName {
	return &FabricFolderInvalidDisplayName{
		resourceType:  "fabric_folder",
		attributeName: "display_name",

		maxLength: 255,
	}
}

// Name returns the rule name
func (r *FabricFolderInvalidDisplayName) Name() string {
	return "fabric_folder_invalid_display_name"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricFolderInvalidDisplayName) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricFolderInvalidDisplayName) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricFolderInvalidDisplayName) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricFolderInvalidDisplayName) Check(runner tflint.Runner) error {
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
