package rules

import (
	"fmt"
	"strings"

	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// FabricDomainContributorsScope validates domain contributors_scope values
type FabricDomainContributorsScope struct {
	tflint.DefaultRule
}

func NewFabricDomainContributorsScope() *FabricDomainContributorsScope {
	return &FabricDomainContributorsScope{}
}

func (r *FabricDomainContributorsScope) Name() string {
	return "fabric_domain_contributors_scope"
}

func (r *FabricDomainContributorsScope) Enabled() bool {
	return true
}

func (r *FabricDomainContributorsScope) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *FabricDomainContributorsScope) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *FabricDomainContributorsScope) Check(runner tflint.Runner) error {
	resourceContent, err := runner.GetResourceContent("fabric_domain", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "contributors_scope"},
		},
	}, nil)
	if err != nil {
		return err
	}

	validScopes := map[string]bool{
		"AdminsOnly":             true,
		"AllTenant":              true,
		"SpecificUsersAndGroups": true,
	}

	for _, resource := range resourceContent.Blocks {
		if attr, exists := resource.Body.Attributes["contributors_scope"]; exists && attr.Expr != nil {
			var scope string
			if err := runner.EvaluateExpr(attr.Expr, &scope, nil); err == nil && scope != "" {
				if !validScopes[scope] {
					validScopesList := []string{"AdminsOnly", "AllTenant", "SpecificUsersAndGroups"}
					if err := runner.EmitIssue(
						r,
						fmt.Sprintf("Invalid contributors_scope '%s'. Must be one of: %s", scope, strings.Join(validScopesList, ", ")),
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
