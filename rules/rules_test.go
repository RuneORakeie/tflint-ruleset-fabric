package rules

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

// TestFabricWorkspaceNaming tests workspace naming rule
func TestFabricWorkspaceNaming(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "valid naming - lowercase with hyphens",
			Content: `
resource "fabric_workspace" "example" {
  display_name = "valid-workspace-name"
  description  = "Test workspace"
  capacity_id  = "capacity-123"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "invalid naming - uppercase letters",
			Content: `
resource "fabric_workspace" "example" {
  display_name = "INVALID-WORKSPACE-NAME"
  description  = "Test workspace"
  capacity_id  = "capacity-123"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewFabricWorkspaceNaming(),
					Message: "Workspace name should be 3-50 characters and contain only lowercase letters, numbers, and hyphens",
				},
			},
		},
		{
			Name: "invalid naming - too short",
			Content: `
resource "fabric_workspace" "example" {
  display_name = "ab"
  description  = "Test workspace"
  capacity_id  = "capacity-123"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewFabricWorkspaceNaming(),
					Message: "Workspace name should be 3-50 characters and contain only lowercase letters, numbers, and hyphens",
				},
			},
		},
	}

	runner := helper.TestRunner(t, cases)
	rule := NewFabricWorkspaceNaming()

	if err := rule.Check(runner); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
}

// TestFabricWorkspaceCapacity tests capacity requirement rule
func TestFabricWorkspaceCapacity(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "valid - capacity assigned",
			Content: `
resource "fabric_workspace" "example" {
  display_name = "test-workspace"
  description  = "Test"
  capacity_id  = "capacity-123"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "invalid - no capacity assigned",
			Content: `
resource "fabric_workspace" "example" {
  display_name = "test-workspace"
  description  = "Test"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewFabricWorkspaceCapacity(),
					Message: "Workspace should have a capacity assigned for production use",
				},
			},
		},
	}

	runner := helper.TestRunner(t, cases)
	rule := NewFabricWorkspaceCapacity()

	if err := rule.Check(runner); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
}

// TestFabricWorkspaceDescription tests description requirement rule
func TestFabricWorkspaceDescription(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "valid - description provided",
			Content: `
resource "fabric_workspace" "example" {
  display_name = "test-workspace"
  description  = "This is a test workspace"
  capacity_id  = "capacity-123"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "invalid - no description",
			Content: `
resource "fabric_workspace" "example" {
  display_name = "test-workspace"
  capacity_id  = "capacity-123"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewFabricWorkspaceDescription(),
					Message: "Workspace should have a description for governance and documentation",
				},
			},
		},
	}

	runner := helper.TestRunner(t, cases)
	rule := NewFabricWorkspaceDescription()

	if err := rule.Check(runner); err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
}
