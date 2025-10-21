package rules

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

// TestFabricWorkspaceNaming tests workspace naming rule
func TestFabricWorkspaceNaming(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "valid naming - lowercase with hyphens",
			content: `
resource "fabric_workspace" "example" {
  display_name = "valid-workspace-name"
  description  = "Test workspace"
  capacity_id  = "capacity-123"
}`,
			hasIssue: false,
		},
		{
			name: "invalid naming - uppercase letters",
			content: `
resource "fabric_workspace" "example" {
  display_name = "INVALID-WORKSPACE-NAME"
  description  = "Test workspace"
  capacity_id  = "capacity-123"
}`,
			hasIssue: true,
		},
		{
			name: "invalid naming - too short",
			content: `
resource "fabric_workspace" "example" {
  display_name = "ab"
  description  = "Test workspace"
  capacity_id  = "capacity-123"
}`,
			hasIssue: true,
		},
		{
			name: "valid naming - exactly 3 characters",
			content: `
resource "fabric_workspace" "example" {
  display_name = "abc"
  description  = "Test workspace"
  capacity_id  = "capacity-123"
}`,
			hasIssue: false,
		},
		{
			name: "invalid naming - too long",
			content: `
resource "fabric_workspace" "example" {
  display_name = "this-is-a-very-long-workspace-name-that-exceeds-fifty-characters-limit"
  description  = "Test workspace"
  capacity_id  = "capacity-123"
}`,
			hasIssue: true,
		},
		{
			name: "invalid naming - special characters",
			content: `
resource "fabric_workspace" "example" {
  display_name = "workspace@123!"
  description  = "Test workspace"
  capacity_id  = "capacity-123"
}`,
			hasIssue: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{
				"main.tf": tt.content,
			})
			rule := NewFabricWorkspaceNaming()

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			if tt.hasIssue && len(runner.Issues) == 0 {
				t.Errorf("Expected issues but got none")
			}
			if !tt.hasIssue && len(runner.Issues) > 0 {
				t.Errorf("Expected no issues but got %d", len(runner.Issues))
			}
		})
	}
}

// TestFabricWorkspaceCapacity tests capacity requirement rule
func TestFabricWorkspaceCapacity(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "valid - capacity assigned",
			content: `
resource "fabric_workspace" "example" {
  display_name = "test-workspace"
  description  = "Test"
  capacity_id  = "capacity-123"
}`,
			hasIssue: false,
		},
		{
			name: "invalid - no capacity assigned",
			content: `
resource "fabric_workspace" "example" {
  display_name = "test-workspace"
  description  = "Test"
}`,
			hasIssue: true,
		},
		{
			name: "valid - capacity assigned with variable",
			content: `
resource "fabric_workspace" "example" {
  display_name = "test-workspace"
  description  = "Test"
  capacity_id  = var.capacity_id
}`,
			hasIssue: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{
				"main.tf": tt.content,
			})
			rule := NewFabricWorkspaceCapacity()

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			if tt.hasIssue && len(runner.Issues) == 0 {
				t.Errorf("Expected issues but got none")
			}
			if !tt.hasIssue && len(runner.Issues) > 0 {
				t.Errorf("Expected no issues but got %d", len(runner.Issues))
			}
		})
	}
}

// TestFabricWorkspaceDescription tests description requirement rule
func TestFabricWorkspaceDescription(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "valid - description provided",
			content: `
resource "fabric_workspace" "example" {
  display_name = "test-workspace"
  description  = "This is a test workspace"
  capacity_id  = "capacity-123"
}`,
			hasIssue: false,
		},
		{
			name: "invalid - no description",
			content: `
resource "fabric_workspace" "example" {
  display_name = "test-workspace"
  capacity_id  = "capacity-123"
}`,
			hasIssue: true,
		},
		{
			name: "valid - description from variable",
			content: `
resource "fabric_workspace" "example" {
  display_name = "test-workspace"
  description  = var.workspace_description
  capacity_id  = "capacity-123"
}`,
			hasIssue: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{
				"main.tf": tt.content,
			})
			rule := NewFabricWorkspaceDescription()

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			if tt.hasIssue && len(runner.Issues) == 0 {
				t.Errorf("Expected issues but got none")
			}
			if !tt.hasIssue && len(runner.Issues) > 0 {
				t.Errorf("Expected no issues but got %d", len(runner.Issues))
			}
		})
	}
}
