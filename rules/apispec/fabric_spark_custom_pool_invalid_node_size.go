package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricSparkCustomPoolInvalidNodeSize checks whether fabric_spark_custom_pool.node_size is valid
type FabricSparkCustomPoolInvalidNodeSize struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string








}

// NewFabricRule returns a new rule instance
func NewFabricSparkCustomPoolInvalidNodeSize() *FabricSparkCustomPoolInvalidNodeSize {
	return &FabricSparkCustomPoolInvalidNodeSize{
		resourceType:  "fabric_spark_custom_pool",
		attributeName: "node_size",
		enum:          []string{ "Small", "Medium", "Large", "XLarge", "XXLarge",  },








	}
}

// Name returns the rule name
func (r *FabricSparkCustomPoolInvalidNodeSize) Name() string {
	return "fabric_spark_custom_pool_invalid_node_size"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricSparkCustomPoolInvalidNodeSize) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricSparkCustomPoolInvalidNodeSize) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricSparkCustomPoolInvalidNodeSize) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricSparkCustomPoolInvalidNodeSize) Check(runner tflint.Runner) error {
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

		if err := r.validateEnum(runner, val, attribute); err != nil {
			return err
		}






	}

	return nil
}

func (r *FabricSparkCustomPoolInvalidNodeSize) validateEnum(runner tflint.Runner, val string, attribute *hclext.Attribute) error {
	for _, valid := range r.enum {
		if val == valid {
			return nil
		}
	}
	return runner.EmitIssue(
		r,
		fmt.Sprintf("node_size must be one of: %v", r.enum),
		attribute.Expr.Range(),
	)
}




