package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hcl"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// FabricWorkspaceNaming checks that workspace names follow naming conventions
type FabricWorkspaceNaming struct {
	tflint.DefaultRule
}

func NewFabricWorkspaceNaming() *FabricWorkspaceNaming {
	return &FabricWorkspaceNaming{}
}

func (r *FabricWorkspaceNaming) Name() string {
	return "fabric_workspace_naming"
}

func (r *FabricWorkspaceNaming) Enabled() bool {
	return true
}

func (r *FabricWorkspaceNaming) Severity() string {
	return tflint.WARNING
}

func (r *FabricWorkspaceNaming) Link() string {
	return "https://learn.microsoft.com/en-us/fabric/admin/fabric-governance"
}

func (r *FabricWorkspaceNaming) Check(runner tflint.Runner) error {
	logger := logger.New()
	logger.Debug("Checking Fabric workspace naming conventions")

	resources, err := runner.GetResourcesByType("fabric_workspace")
	if err != nil {
		return err
	}

	namePattern := regexp.MustCompile(`^[a-z0-9\-]{3,50}$`)

	for _, resource := range resources {
		var name string
		err := resource.GetAttribute("display_name", &name)
		if err == nil && name != "" {
			if !namePattern.MatchString(name) {
				runner.EmitIssue(
					r,
					"Workspace name should be 3-50 characters and contain only lowercase letters, numbers, and hyphens",
					resource.GetNameRange(),
				)
			}
		}
	}

	return nil
}