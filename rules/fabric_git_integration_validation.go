package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// FabricGitIntegrationValidation validates Git integration configuration
type FabricGitIntegrationValidation struct {
	tflint.DefaultRule
}

func NewFabricGitIntegrationValidation() *FabricGitIntegrationValidation {
	return &FabricGitIntegrationValidation{}
}

func (r *FabricGitIntegrationValidation) Name() string {
	return "fabric_git_integration_provider_valid"
}

func (r *FabricGitIntegrationValidation) Enabled() bool {
	return true
}

func (r *FabricGitIntegrationValidation) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *FabricGitIntegrationValidation) Link() string {
	return "https://learn.microsoft.com/en-us/fabric/cicd/git-integration/intro-to-git-integration"
}

func (r *FabricGitIntegrationValidation) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("fabric_workspace_git_connection", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "git_provider_type"},
		},
	}, nil)
	if err != nil {
		return err
	}

	validProviders := map[string]bool{
		"GitHub":           true,
		"Azure DevOps":     true,
		"Bitbucket Cloud":  true,
		"GitLab":           true,
	}

	if attr, exists := resources.Attributes["git_provider_type"]; exists && attr.Expr != nil {
		var provider string
		if err := runner.EvaluateExpr(attr.Expr, &provider, nil); err == nil && provider != "" {
			if !validProviders[provider] {
				runner.EmitIssue(
					r,
					fmt.Sprintf("Invalid Git provider: %s. Must be one of: GitHub, Azure DevOps, Bitbucket Cloud, GitLab", provider),
					attr.Range,
				)
			}
		}
	}

	return nil
}
