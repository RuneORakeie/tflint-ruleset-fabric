package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hcl"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// FabricWorkspaceCapacity ensures capacity is assigned to workspaces
type FabricWorkspaceCapacity struct {
	tflint.DefaultRule
}

func NewFabricWorkspaceCapacity() *FabricWorkspaceCapacity {
	return &FabricWorkspaceCapacity{}
}

func (r *FabricWorkspaceCapacity) Name() string {
	return "fabric_workspace_capacity_required"
}

func (r *FabricWorkspaceCapacity) Enabled() bool {
	return true
}

func (r *FabricWorkspaceCapacity) Severity() string {
	return tflint.ERROR
}

func (r *FabricWorkspaceCapacity) Link() string {
	return "https://learn.microsoft.com/en-us/fabric/admin/capacity-settings"
}

func (r *FabricWorkspaceCapacity) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourcesByType("fabric_workspace")
	if err != nil {
		return err
	}

	for _, resource := range resources {
		var capacity string
		err := resource.GetAttribute("capacity_id", &capacity)
		if err != nil || capacity == "" {
			runner.EmitIssue(
				r,
				"Workspace should have a capacity assigned for production use",
				resource.GetNameRange(),
			)
		}
	}

	return nil
}