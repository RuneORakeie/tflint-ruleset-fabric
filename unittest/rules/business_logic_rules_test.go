package rules

import (
	"testing"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

// TestFabricCapacityRegion tests region validation rule
func TestFabricCapacityRegion(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "valid region - westeurope",
			content: `resource "fabric_capacity" "example" {
				region = "westeurope"
			}`,
			hasIssue: false,
		},
		{
			name: "valid region - eastus",
			content: `resource "fabric_capacity" "example" {
				region = "eastus"
			}`,
			hasIssue: false,
		},
		{
			name: "invalid region",
			content: `resource "fabric_capacity" "example" {
				region = "invalid-region"
			}`,
			hasIssue: true,
		},
		{
			name: "no region attribute",
			content: `resource "fabric_capacity" "example" {
			}`,
			hasIssue: false,
		},
	}
	
	rule := NewFabricCapacityRegion()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"main.tf": tt.content})
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if tt.hasIssue {
				helper.AssertIssuesCount(t, 1, runner.Issues)
			} else {
				helper.AssertIssuesCount(t, 0, runner.Issues)
			}
		})
	}
}

// TestFabricDeploymentPipelineStagesCount tests stages count validation
func TestFabricDeploymentPipelineStagesCount(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "valid - 2 stages",
			content: `resource "fabric_deployment_pipeline" "example" {
				display_name = "Test"
				stages {
					display_name = "dev"
					is_public = true
				}
				stages {
					display_name = "prod"
					is_public = true
				}
			}`,
			hasIssue: false,
		},
		{
			name: "valid - 10 stages (max)",
			content: `resource "fabric_deployment_pipeline" "example" {
				display_name = "Test"
				stages {
					display_name = "stage1"
					is_public = true
				}
				stages {
					display_name = "stage2"
					is_public = true
				}
				stages {
					display_name = "stage3"
					is_public = true
				}
				stages {
					display_name = "stage4"
					is_public = true
				}
				stages {
					display_name = "stage5"
					is_public = true
				}
				stages {
					display_name = "stage6"
					is_public = true
				}
				stages {
					display_name = "stage7"
					is_public = true
				}
				stages {
					display_name = "stage8"
					is_public = true
				}
				stages {
					display_name = "stage9"
					is_public = true
				}
				stages {
					display_name = "stage10"
					is_public = true
				}
			}`,
			hasIssue: false,
		},
		{
			name: "invalid - 1 stage (minimum is 2)",
			content: `resource "fabric_deployment_pipeline" "example" {
				display_name = "Test"
				stages {
					display_name = "dev"
					is_public = true
				}
			}`,
			hasIssue: true,
		},
		{
			name: "invalid - 11 stages (exceeds max)",
			content: `resource "fabric_deployment_pipeline" "example" {
				display_name = "Test"
				stages {
					display_name = "stage1"
					is_public = true
				}
				stages {
					display_name = "stage2"
					is_public = true
				}
				stages {
					display_name = "stage3"
					is_public = true
				}
				stages {
					display_name = "stage4"
					is_public = true
				}
				stages {
					display_name = "stage5"
					is_public = true
				}
				stages {
					display_name = "stage6"
					is_public = true
				}
				stages {
					display_name = "stage7"
					is_public = true
				}
				stages {
					display_name = "stage8"
					is_public = true
				}
				stages {
					display_name = "stage9"
					is_public = true
				}
				stages {
					display_name = "stage10"
					is_public = true
				}
				stages {
					display_name = "stage11"
					is_public = true
				}
			}`,
			hasIssue: true,
		},
	}
	
	rule := NewFabricDeploymentPipelineStagesCount()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"main.tf": tt.content})
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if tt.hasIssue {
				helper.AssertIssuesCount(t, 1, runner.Issues)
			} else {
				helper.AssertIssuesCount(t, 0, runner.Issues)
			}
		})
	}
}

// TestFabricDeploymentPipelineStagesDescriptionLength tests stage description length
func TestFabricDeploymentPipelineStagesDescriptionLength(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "valid description",
			content: `resource "fabric_deployment_pipeline" "example" {
				display_name = "Test"
				stages {
					display_name = "dev"
					description = "Development environment"
					is_public = true
				}
				stages {
					display_name = "prod"
					is_public = true
				}
			}`,
			hasIssue: false,
		},
		{
			name: "valid - at max length (256)",
			content: `resource "fabric_deployment_pipeline" "example" {
				display_name = "Test"
				stages {
					display_name = "dev"
					description = "` + string(make([]byte, 256)) + `"
					is_public = true
				}
				stages {
					display_name = "prod"
					is_public = true
				}
			}`,
			hasIssue: false,
		},
		{
			name: "invalid - exceeds max length",
			content: `resource "fabric_deployment_pipeline" "example" {
				display_name = "Test"
				stages {
					display_name = "dev"
					description = "` + string(make([]byte, 257)) + `"
					is_public = true
				}
				stages {
					display_name = "prod"
					is_public = true
				}
			}`,
			hasIssue: true,
		},
	}
	
	rule := NewFabricDeploymentPipelineStagesDescriptionLength()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"main.tf": tt.content})
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if tt.hasIssue {
				helper.AssertIssuesCount(t, 1, runner.Issues)
			} else {
				helper.AssertIssuesCount(t, 0, runner.Issues)
			}
		})
	}
}

// TestFabricDeploymentPipelineStagesDisplayNameLength tests stage display name length
func TestFabricDeploymentPipelineStagesDisplayNameLength(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "valid display name",
			content: `resource "fabric_deployment_pipeline" "example" {
				display_name = "Test"
				stages {
					display_name = "Development"
					is_public = true
				}
				stages {
					display_name = "Production"
					is_public = true
				}
			}`,
			hasIssue: false,
		},
		{
			name: "valid - at max length (256)",
			content: `resource "fabric_deployment_pipeline" "example" {
				display_name = "Test"
				stages {
					display_name = "` + string(make([]byte, 256)) + `"
					is_public = true
				}
				stages {
					display_name = "prod"
					is_public = true
				}
			}`,
			hasIssue: false,
		},
		{
			name: "invalid - exceeds max length",
			content: `resource "fabric_deployment_pipeline" "example" {
				display_name = "Test"
				stages {
					display_name = "` + string(make([]byte, 257)) + `"
					is_public = true
				}
				stages {
					display_name = "prod"
					is_public = true
				}
			}`,
			hasIssue: true,
		},
	}
	
	rule := NewFabricDeploymentPipelineStagesDisplayNameLength()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"main.tf": tt.content})
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if tt.hasIssue {
				helper.AssertIssuesCount(t, 1, runner.Issues)
			} else {
				helper.AssertIssuesCount(t, 0, runner.Issues)
			}
		})
	}
}

// TestFabricDomainContributorsScope tests contributors scope validation
func TestFabricDomainContributorsScope(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "valid - Workspace",
			content: `resource "fabric_domain" "example" {
				display_name = "Test Domain"
				contributors_scope = "Workspace"
			}`,
			hasIssue: false,
		},
		{
			name: "valid - DomainOnly",
			content: `resource "fabric_domain" "example" {
				display_name = "Test Domain"
				contributors_scope = "DomainOnly"
			}`,
			hasIssue: false,
		},
		{
			name: "valid - SpecificWorkspaces",
			content: `resource "fabric_domain" "example" {
				display_name = "Test Domain"
				contributors_scope = "SpecificWorkspaces"
			}`,
			hasIssue: false,
		},
		{
			name: "invalid scope",
			content: `resource "fabric_domain" "example" {
				display_name = "Test Domain"
				contributors_scope = "InvalidScope"
			}`,
			hasIssue: true,
		},
	}
	
	rule := NewFabricDomainContributorsScope()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"main.tf": tt.content})
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if tt.hasIssue {
				helper.AssertIssuesCount(t, 1, runner.Issues)
			} else {
				helper.AssertIssuesCount(t, 0, runner.Issues)
			}
		})
	}
}

// TestFabricItemDescriptionRecommended tests description recommendations
func TestFabricItemDescriptionRecommended(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "valid - workspace with description",
			content: `resource "fabric_workspace" "example" {
				display_name = "Test"
				description = "Test workspace"
			}`,
			hasIssue: false,
		},
		{
			name: "valid - lakehouse with description",
			content: `resource "fabric_lakehouse" "example" {
				workspace_id = "test"
				display_name = "Test"
				description = "Test lakehouse"
			}`,
			hasIssue: false,
		},
		{
			name: "warning - workspace without description",
			content: `resource "fabric_workspace" "example" {
				display_name = "Test"
			}`,
			hasIssue: true,
		},
		{
			name: "warning - empty description",
			content: `resource "fabric_workspace" "example" {
				display_name = "Test"
				description = ""
			}`,
			hasIssue: true,
		},
	}
	
	rule := NewFabricItemDescriptionRecommended()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"main.tf": tt.content})
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if tt.hasIssue {
				helper.AssertIssuesCount(t, 1, runner.Issues)
			} else {
				helper.AssertIssuesCount(t, 0, runner.Issues)
			}
		})
	}
}

// TestFabricRoleAssignmentRecommended tests role assignment recommendations
func TestFabricRoleAssignmentRecommended(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "valid - workspace with role assignment",
			content: `
resource "fabric_workspace" "example" {
	display_name = "Test"
}

resource "fabric_workspace_role_assignment" "example" {
	workspace_id = fabric_workspace.example.id
	principal_id = "test-id"
	principal_type = "User"
	role = "Admin"
}`,
			hasIssue: false,
		},
		{
			name: "warning - workspace without role assignment",
			content: `resource "fabric_workspace" "example" {
				display_name = "Test"
			}`,
			hasIssue: true,
		},
		{
			name: "valid - deployment pipeline with role assignment",
			content: `
resource "fabric_deployment_pipeline" "example" {
	display_name = "Test"
	stages {
		display_name = "dev"
		is_public = true
	}
	stages {
		display_name = "prod"
		is_public = true
	}
}

resource "fabric_deployment_pipeline_role_assignment" "example" {
	deployment_pipeline_id = fabric_deployment_pipeline.example.id
	principal_id = "test-id"
	principal_type = "User"
	role = "Admin"
}`,
			hasIssue: false,
		},
	}
	
	rule := NewFabricRoleAssignmentRecommended()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"main.tf": tt.content})
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if tt.hasIssue {
				helper.AssertIssuesCount(t, 1, runner.Issues)
			} else {
				helper.AssertIssuesCount(t, 0, runner.Issues)
			}
		})
	}
}

// TestFabricWorkspaceCapacity tests capacity requirement
func TestFabricWorkspaceCapacity(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "valid - with capacity",
			content: `resource "fabric_workspace" "example" {
				display_name = "Test"
				capacity_id = "test-capacity-id"
			}`,
			hasIssue: false,
		},
		{
			name: "error - without capacity",
			content: `resource "fabric_workspace" "example" {
				display_name = "Test"
			}`,
			hasIssue: true,
		},
	}
	
	rule := NewFabricWorkspaceCapacity()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"main.tf": tt.content})
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if tt.hasIssue {
				helper.AssertIssuesCount(t, 1, runner.Issues)
			} else {
				helper.AssertIssuesCount(t, 0, runner.Issues)
			}
		})
	}
}

// TestFabricWorkspaceDescription tests workspace description requirement
func TestFabricWorkspaceDescription(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "valid - with description",
			content: `resource "fabric_workspace" "example" {
				display_name = "Test"
				description = "Test workspace"
			}`,
			hasIssue: false,
		},
		{
			name: "warning - without description",
			content: `resource "fabric_workspace" "example" {
				display_name = "Test"
			}`,
			hasIssue: true,
		},
	}
	
	rule := NewFabricWorkspaceDescription()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"main.tf": tt.content})
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if tt.hasIssue {
				helper.AssertIssuesCount(t, 1, runner.Issues)
			} else {
				helper.AssertIssuesCount(t, 0, runner.Issues)
			}
		})
	}
}

// TestFabricWorkspaceGitAzdoAttributes tests Azure DevOps git attributes
func TestFabricWorkspaceGitAzdoAttributes(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "valid - Azure DevOps with all required attributes",
			content: `resource "fabric_workspace_git" "example" {
				workspace_id = "test"
				initialization_strategy = "PreferWorkspace"
				git_provider_details {
					git_provider_type = "AzureDevOps"
					organization_name = "myorg"
					project_name = "myproject"
					repository_name = "myrepo"
					branch_name = "main"
					directory_name = "/test"
				}
			}`,
			hasIssue: false,
		},
		{
			name: "error - Azure DevOps missing organization_name",
			content: `resource "fabric_workspace_git" "example" {
				workspace_id = "test"
				initialization_strategy = "PreferWorkspace"
				git_provider_details {
					git_provider_type = "AzureDevOps"
					project_name = "myproject"
					repository_name = "myrepo"
					branch_name = "main"
					directory_name = "/test"
				}
			}`,
			hasIssue: true,
		},
		{
			name: "error - Azure DevOps missing project_name",
			content: `resource "fabric_workspace_git" "example" {
				workspace_id = "test"
				initialization_strategy = "PreferWorkspace"
				git_provider_details {
					git_provider_type = "AzureDevOps"
					organization_name = "myorg"
					repository_name = "myrepo"
					branch_name = "main"
					directory_name = "/test"
				}
			}`,
			hasIssue: true,
		},
		{
			name: "valid - GitHub (should not check Azure DevOps attributes)",
			content: `resource "fabric_workspace_git" "example" {
				workspace_id = "test"
				initialization_strategy = "PreferWorkspace"
				git_provider_details {
					git_provider_type = "GitHub"
					repository_name = "myrepo"
					branch_name = "main"
					directory_name = "/test"
				}
			}`,
			hasIssue: false,
		},
	}
	
	rule := NewFabricWorkspaceGitAzdoAttributes()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"main.tf": tt.content})
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if tt.hasIssue {
				helper.AssertIssuesCount(t, 1, runner.Issues)
			} else {
				helper.AssertIssuesCount(t, 0, runner.Issues)
			}
		})
	}
}

// TestFabricWorkspaceGitCredentialsSource tests git credentials source validation
func TestFabricWorkspaceGitCredentialsSource(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "valid - ServicePrincipal",
			content: `resource "fabric_workspace_git" "example" {
				workspace_id = "test"
				initialization_strategy = "PreferWorkspace"
				credentials_source = "ServicePrincipal"
				git_provider_details {
					git_provider_type = "GitHub"
					repository_name = "myrepo"
					branch_name = "main"
					directory_name = "/test"
				}
			}`,
			hasIssue: false,
		},
		{
			name: "valid - UserAuthentication",
			content: `resource "fabric_workspace_git" "example" {
				workspace_id = "test"
				initialization_strategy = "PreferWorkspace"
				credentials_source = "UserAuthentication"
				git_provider_details {
					git_provider_type = "GitHub"
					repository_name = "myrepo"
					branch_name = "main"
					directory_name = "/test"
				}
			}`,
			hasIssue: false,
		},
		{
			name: "error - invalid credentials source",
			content: `resource "fabric_workspace_git" "example" {
				workspace_id = "test"
				initialization_strategy = "PreferWorkspace"
				credentials_source = "InvalidSource"
				git_provider_details {
					git_provider_type = "GitHub"
					repository_name = "myrepo"
					branch_name = "main"
					directory_name = "/test"
				}
			}`,
			hasIssue: true,
		},
	}
	
	rule := NewFabricWorkspaceGitCredentialsSource()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"main.tf": tt.content})
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if tt.hasIssue {
				helper.AssertIssuesCount(t, 1, runner.Issues)
			} else {
				helper.AssertIssuesCount(t, 0, runner.Issues)
			}
		})
	}
}

// TestFabricWorkspaceGitDirectoryName tests directory name validation
func TestFabricWorkspaceGitDirectoryName(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "valid - starts with /",
			content: `resource "fabric_workspace_git" "example" {
				workspace_id = "test"
				initialization_strategy = "PreferWorkspace"
				git_provider_details {
					git_provider_type = "GitHub"
					repository_name = "myrepo"
					branch_name = "main"
					directory_name = "/test"
				}
			}`,
			hasIssue: false,
		},
		{
			name: "error - doesn't start with /",
			content: `resource "fabric_workspace_git" "example" {
				workspace_id = "test"
				initialization_strategy = "PreferWorkspace"
				git_provider_details {
					git_provider_type = "GitHub"
					repository_name = "myrepo"
					branch_name = "main"
					directory_name = "test"
				}
			}`,
			hasIssue: true,
		},
		{
			name: "error - exceeds max length",
			content: `resource "fabric_workspace_git" "example" {
				workspace_id = "test"
				initialization_strategy = "PreferWorkspace"
				git_provider_details {
					git_provider_type = "GitHub"
					repository_name = "myrepo"
					branch_name = "main"
					directory_name = "/` + string(make([]byte, 256)) + `"
				}
			}`,
			hasIssue: true,
		},
	}
	
	rule := NewFabricWorkspaceGitDirectoryName()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"main.tf": tt.content})
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if tt.hasIssue {
				helper.AssertIssuesCount(t, 1, runner.Issues)
			} else {
				helper.AssertIssuesCount(t, 0, runner.Issues)
			}
		})
	}
}

// TestFabricWorkspaceGitGithubAttributes tests GitHub git attributes
func TestFabricWorkspaceGitGithubAttributes(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "valid - GitHub with all required attributes",
			content: `resource "fabric_workspace_git" "example" {
				workspace_id = "test"
				initialization_strategy = "PreferWorkspace"
				git_provider_details {
					git_provider_type = "GitHub"
					owner_name = "myowner"
					repository_name = "myrepo"
					branch_name = "main"
					directory_name = "/test"
				}
			}`,
			hasIssue: false,
		},
		{
			name: "error - GitHub missing owner_name",
			content: `resource "fabric_workspace_git" "example" {
				workspace_id = "test"
				initialization_strategy = "PreferWorkspace"
				git_provider_details {
					git_provider_type = "GitHub"
					repository_name = "myrepo"
					branch_name = "main"
					directory_name = "/test"
				}
			}`,
			hasIssue: true,
		},
		{
			name: "valid - Azure DevOps (should not check GitHub attributes)",
			content: `resource "fabric_workspace_git" "example" {
				workspace_id = "test"
				initialization_strategy = "PreferWorkspace"
				git_provider_details {
					git_provider_type = "AzureDevOps"
					organization_name = "myorg"
					project_name = "myproject"
					repository_name = "myrepo"
					branch_name = "main"
					directory_name = "/test"
				}
			}`,
			hasIssue: false,
		},
	}
	
	rule := NewFabricWorkspaceGitGithubAttributes()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"main.tf": tt.content})
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if tt.hasIssue {
				helper.AssertIssuesCount(t, 1, runner.Issues)
			} else {
				helper.AssertIssuesCount(t, 0, runner.Issues)
			}
		})
	}
}

// TestFabricWorkspaceGitInitializationStrategy tests initialization strategy validation
func TestFabricWorkspaceGitInitializationStrategy(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "valid - PreferWorkspace",
			content: `resource "fabric_workspace_git" "example" {
				workspace_id = "test"
				initialization_strategy = "PreferWorkspace"
				git_provider_details {
					git_provider_type = "GitHub"
					repository_name = "myrepo"
					branch_name = "main"
					directory_name = "/test"
				}
			}`,
			hasIssue: false,
		},
		{
			name: "valid - PreferRemote",
			content: `resource "fabric_workspace_git" "example" {
				workspace_id = "test"
				initialization_strategy = "PreferRemote"
				git_provider_details {
					git_provider_type = "GitHub"
					repository_name = "myrepo"
					branch_name = "main"
					directory_name = "/test"
				}
			}`,
			hasIssue: false,
		},
		{
			name: "error - invalid strategy",
			content: `resource "fabric_workspace_git" "example" {
				workspace_id = "test"
				initialization_strategy = "InvalidStrategy"
				git_provider_details {
					git_provider_type = "GitHub"
					repository_name = "myrepo"
					branch_name = "main"
					directory_name = "/test"
				}
			}`,
			hasIssue: true,
		},
	}
	
	rule := NewFabricWorkspaceGitInitializationStrategy()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"main.tf": tt.content})
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if tt.hasIssue {
				helper.AssertIssuesCount(t, 1, runner.Issues)
			} else {
				helper.AssertIssuesCount(t, 0, runner.Issues)
			}
		})
	}
}

// TestFabricWorkspaceGitProviderType tests git provider type validation
func TestFabricWorkspaceGitProviderType(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "valid - GitHub",
			content: `resource "fabric_workspace_git" "example" {
				workspace_id = "test"
				initialization_strategy = "PreferWorkspace"
				git_provider_details {
					git_provider_type = "GitHub"
					repository_name = "myrepo"
					branch_name = "main"
					directory_name = "/test"
				}
			}`,
			hasIssue: false,
		},
		{
			name: "valid - AzureDevOps",
			content: `resource "fabric_workspace_git" "example" {
				workspace_id = "test"
				initialization_strategy = "PreferWorkspace"
				git_provider_details {
					git_provider_type = "AzureDevOps"
					organization_name = "myorg"
					project_name = "myproject"
					repository_name = "myrepo"
					branch_name = "main"
					directory_name = "/test"
				}
			}`,
			hasIssue: false,
		},
		{
			name: "error - invalid provider type",
			content: `resource "fabric_workspace_git" "example" {
				workspace_id = "test"
				initialization_strategy = "PreferWorkspace"
				git_provider_details {
					git_provider_type = "GitLab"
					repository_name = "myrepo"
					branch_name = "main"
					directory_name = "/test"
				}
			}`,
			hasIssue: true,
		},
	}
	
	rule := NewFabricWorkspaceGitProviderType()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"main.tf": tt.content})
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if tt.hasIssue {
				helper.AssertIssuesCount(t, 1, runner.Issues)
			} else {
				helper.AssertIssuesCount(t, 0, runner.Issues)
			}
		})
	}
}

// TestFabricWorkspaceGitStringLengths tests git string length validations
func TestFabricWorkspaceGitStringLengths(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "valid - all within limits",
			content: `resource "fabric_workspace_git" "example" {
				workspace_id = "test"
				initialization_strategy = "PreferWorkspace"
				git_provider_details {
					git_provider_type = "GitHub"
					owner_name = "owner"
					repository_name = "repo"
					branch_name = "main"
					directory_name = "/test"
				}
			}`,
			hasIssue: false,
		},
		{
			name: "error - repository_name exceeds limit",
			content: `resource "fabric_workspace_git" "example" {
				workspace_id = "test"
				initialization_strategy = "PreferWorkspace"
				git_provider_details {
					git_provider_type = "GitHub"
					repository_name = "` + string(make([]byte, 101)) + `"
					branch_name = "main"
					directory_name = "/test"
				}
			}`,
			hasIssue: true,
		},
	}
	
	rule := NewFabricWorkspaceGitStringLengths()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"main.tf": tt.content})
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if tt.hasIssue {
				helper.AssertIssuesCount(t, 1, runner.Issues)
			} else {
				helper.AssertIssuesCount(t, 0, runner.Issues)
			}
		})
	}
}

// TestFabricWorkspaceRoleAssignmentRole tests workspace role validation
func TestFabricWorkspaceRoleAssignmentRole(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "valid - Admin role",
			content: `resource "fabric_workspace_role_assignment" "example" {
				workspace_id = "test"
				principal_id = "test-id"
				principal_type = "User"
				role = "Admin"
			}`,
			hasIssue: false,
		},
		{
			name: "valid - Member role",
			content: `resource "fabric_workspace_role_assignment" "example" {
				workspace_id = "test"
				principal_id = "test-id"
				principal_type = "User"
				role = "Member"
			}`,
			hasIssue: false,
		},
		{
			name: "valid - Contributor role",
			content: `resource "fabric_workspace_role_assignment" "example" {
				workspace_id = "test"
				principal_id = "test-id"
				principal_type = "User"
				role = "Contributor"
			}`,
			hasIssue: false,
		},
		{
			name: "valid - Viewer role",
			content: `resource "fabric_workspace_role_assignment" "example" {
				workspace_id = "test"
				principal_id = "test-id"
				principal_type = "User"
				role = "Viewer"
			}`,
			hasIssue: false,
		},
		{
			name: "error - invalid role",
			content: `resource "fabric_workspace_role_assignment" "example" {
				workspace_id = "test"
				principal_id = "test-id"
				principal_type = "User"
				role = "InvalidRole"
			}`,
			hasIssue: true,
		},
	}
	
	rule := NewFabricWorkspaceRoleAssignmentRole()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"main.tf": tt.content})
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if tt.hasIssue {
				helper.AssertIssuesCount(t, 1, runner.Issues)
			} else {
				helper.AssertIssuesCount(t, 0, runner.Issues)
			}
		})
	}
}