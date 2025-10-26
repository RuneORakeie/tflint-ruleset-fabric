package apispec

import (
	"fmt"
	"regexp"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricEventstreamInvalidDisplayName checks whether fabric_eventstream.display_name is valid
type FabricEventstreamInvalidDisplayName struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string

	pattern       string



	maxLength     int



}

// NewFabricRule returns a new rule instance
func NewFabricEventstreamInvalidDisplayName() *FabricEventstreamInvalidDisplayName {
	return &FabricEventstreamInvalidDisplayName{
		resourceType:  "fabric_eventstream",
		attributeName: "display_name",

		pattern:       `^[a-zA-Z0-9._-]+$`,



		maxLength:     256,



	}
}

// Name returns the rule name
func (r *FabricEventstreamInvalidDisplayName) Name() string {
	return "fabric_eventstream_invalid_display_name"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricEventstreamInvalidDisplayName) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricEventstreamInvalidDisplayName) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricEventstreamInvalidDisplayName) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricEventstreamInvalidDisplayName) Check(runner tflint.Runner) error {
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


func (r *FabricEventstreamInvalidDisplayName) validatePattern(runner tflint.Runner, val string, attribute *hclext.Attribute) error {
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



