package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hcl"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// FabricWorkspaceDescription ensures workspaces have descriptions
type FabricWorkspaceDescription struct {
	tflint.DefaultRule
}

func NewFabricWorkspaceDescription() *FabricWorkspaceDescription {
	return &FabricWorkspaceDescription{}
}

func (r *FabricWorkspaceDescription) Name() string {
	return "fabric_workspace_description_required"
}

func (r *FabricWorkspaceDescription) Enabled() bool {
	return true
}

func (r *FabricWorkspaceDescription) Severity() string {
	return tflint.WARNING
}

func (r *FabricWorkspaceDescription) Link() string {
	return "https://learn.microsoft.com/en-us/fabric/admin/fabric-governance"
}

func (r *FabricWorkspaceDescription) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourcesByType("fabric_workspace")
	if err != nil {
		return err
	}

	for _, resource := range resources {
		var description string
		err := resource.GetAttribute("description", &description)
		if err != nil || description == "" {
			runner.EmitIssue(
				r,
				"Workspace should have a description for governance and documentation",
				resource.GetNameRange(),
			)
		}
	}

	return nil
}

