package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hcl"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
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

func (r *FabricGitIntegrationValidation) Severity() string {
	return tflint.ERROR
}

func (r *FabricGitIntegrationValidation) Link() string {
	return "https://learn.microsoft.com/en-us/fabric/cicd/git-integration/intro-to-git-integration"
}

func (r *FabricGitIntegrationValidation) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourcesByType("fabric_workspace_git_connection")
	if err != nil {
		return err
	}

	validProviders := map[string]bool{
		"GitHub":           true,
		"Azure DevOps":     true,
	}

	for _, resource := range resources {
		var provider string
		err := resource.GetAttribute("git_provider_type", &provider)
		if err == nil && provider != "" {
			if !validProviders[provider] {
				runner.EmitIssue(
					r,
					fmt.Sprintf("Invalid Git provider: %s. Must be one of: GitHub or Azure DevOps", provider),
					resource.GetNameRange(),
				)
			}
		}
	}

	return nil
}
