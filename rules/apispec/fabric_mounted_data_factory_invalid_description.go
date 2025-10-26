package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricMountedDataFactoryInvalidDescription checks whether fabric_mounted_data_factory.description is valid
type FabricMountedDataFactoryInvalidDescription struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string





	maxLength     int



}

// NewFabricRule returns a new rule instance
func NewFabricMountedDataFactoryInvalidDescription() *FabricMountedDataFactoryInvalidDescription {
	return &FabricMountedDataFactoryInvalidDescription{
		resourceType:  "fabric_mounted_data_factory",
		attributeName: "description",





		maxLength:     256,



	}
}

// Name returns the rule name
func (r *FabricMountedDataFactoryInvalidDescription) Name() string {
	return "fabric_mounted_data_factory_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricMountedDataFactoryInvalidDescription) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricMountedDataFactoryInvalidDescription) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricMountedDataFactoryInvalidDescription) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricMountedDataFactoryInvalidDescription) Check(runner tflint.Runner) error {
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
				fmt.Sprintf("description must be at most %d characters (actual: %d)", r.maxLength, len(val)),
				attribute.Expr.Range(),
			)
		}



	}

	return nil
}




