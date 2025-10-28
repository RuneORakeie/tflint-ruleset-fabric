package apispec

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)


// FabricEventhouseInvalidDisplayName checks whether fabric_eventhouse.display_name is valid
type FabricEventhouseInvalidDisplayName struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string

	pattern string

	maxLength int
}

// NewFabricRule returns a new rule instance
func NewFabricEventhouseInvalidDisplayName() *FabricEventhouseInvalidDisplayName {
	return &FabricEventhouseInvalidDisplayName{
		resourceType:  "fabric_eventhouse",
		attributeName: "display_name",

		pattern: `^[a-zA-Z0-9._-]+$`,

		maxLength: 256,
	}
}

// Name returns the rule name
func (r *FabricEventhouseInvalidDisplayName) Name() string {
	return "fabric_eventhouse_invalid_display_name"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricEventhouseInvalidDisplayName) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricEventhouseInvalidDisplayName) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricEventhouseInvalidDisplayName) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricEventhouseInvalidDisplayName) Check(runner tflint.Runner) error {
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

		if err := r.validatePattern(runner, val, attribute); err != nil {
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

func (r *FabricEventhouseInvalidDisplayName) validatePattern(runner tflint.Runner, val string, attribute *hclext.Attribute) error {
	matched, err := regexp.MatchString(r.pattern, val)
	if err != nil {
		return err
	}
	if !matched {
		return runner.EmitIssue(
			r,
			fmt.Sprintf("display_name does not match required pattern: %s", r.pattern),
			attribute.Expr.Range(),
		)
	}
	return nil
}
