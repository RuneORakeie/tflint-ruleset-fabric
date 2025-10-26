package apispec

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricConnectionInvalidPrivacyLevel checks whether fabric_connection.privacy_level is valid
type FabricConnectionInvalidPrivacyLevel struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string








}

// NewFabricRule returns a new rule instance
func NewFabricConnectionInvalidPrivacyLevel() *FabricConnectionInvalidPrivacyLevel {
	return &FabricConnectionInvalidPrivacyLevel{
		resourceType:  "fabric_connection",
		attributeName: "privacy_level",
		enum:          []string{ "None", "Private", "Organizational", "Public",  },








	}
}

// Name returns the rule name
func (r *FabricConnectionInvalidPrivacyLevel) Name() string {
	return "fabric_connection_invalid_privacy_level"
}

// Enabled returns whether the rule is enabled by default
func (r *FabricConnectionInvalidPrivacyLevel) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FabricConnectionInvalidPrivacyLevel) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FabricConnectionInvalidPrivacyLevel) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check validates the resource
func (r *FabricConnectionInvalidPrivacyLevel) Check(runner tflint.Runner) error {
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

func (r *FabricConnectionInvalidPrivacyLevel) validateEnum(runner tflint.Runner, val string, attribute *hclext.Attribute) error {
	for _, valid := range r.enum {
		if val == valid {
			return nil
		}
	}
	return runner.EmitIssue(
		r,
		fmt.Sprintf("privacy_level must be one of: %v", r.enum),
		attribute.Expr.Range(),
	)
}




