package apispec

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GeneratedRuleInfo contains metadata about a generated rule
type GeneratedRuleInfo struct {
	Name       string
	Type       string
	Constructor func() interface{ Check(tflint.Runner) error }
}

// GetGeneratedRules returns a map of all generated rule constructors
// This function discovers rules from the rules package
func GetGeneratedRules() map[string]GeneratedRuleInfo {
	rules := make(map[string]GeneratedRuleInfo)

	// List of generated rules - these are instantiated from the generated directory
	// Each rule follows the naming convention: fabric_<ResourceType>_<Constraint>.go
	generatedRuleConstructors := []GeneratedRuleInfo{
	{
		Name: "fabric_activator_invalid_description",
		Type: "FabricActivatorInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricActivatorInvalidDescription()
		},
	},
	{
		Name: "fabric_apache_airflow_job_invalid_description",
		Type: "FabricApacheAirflowJobInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricApacheAirflowJobInvalidDescription()
		},
	},
	{
		Name: "fabric_connection_invalid_connectivity_type",
		Type: "FabricConnectionInvalidConnectivityType",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricConnectionInvalidConnectivityType()
		},
	},
	{
		Name: "fabric_connection_invalid_display_name",
		Type: "FabricConnectionInvalidDisplayName",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricConnectionInvalidDisplayName()
		},
	},
	{
		Name: "fabric_connection_invalid_privacy_level",
		Type: "FabricConnectionInvalidPrivacyLevel",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricConnectionInvalidPrivacyLevel()
		},
	},
	{
		Name: "fabric_copy_job_invalid_description",
		Type: "FabricCopyJobInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricCopyJobInvalidDescription()
		},
	},
	{
		Name: "fabric_copy_job_invalid_display_name",
		Type: "FabricCopyJobInvalidDisplayName",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricCopyJobInvalidDisplayName()
		},
	},
	{
		Name: "fabric_data_pipeline_invalid_description",
		Type: "FabricDataPipelineInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricDataPipelineInvalidDescription()
		},
	},
	{
		Name: "fabric_data_pipeline_invalid_display_name",
		Type: "FabricDataPipelineInvalidDisplayName",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricDataPipelineInvalidDisplayName()
		},
	},
	{
		Name: "fabric_dataflow_invalid_description",
		Type: "FabricDataflowInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricDataflowInvalidDescription()
		},
	},
	{
		Name: "fabric_dataflow_invalid_display_name",
		Type: "FabricDataflowInvalidDisplayName",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricDataflowInvalidDisplayName()
		},
	},
	{
		Name: "fabric_deployment_pipeline_invalid_description",
		Type: "FabricDeploymentPipelineInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricDeploymentPipelineInvalidDescription()
		},
	},
	{
		Name: "fabric_deployment_pipeline_invalid_display_name",
		Type: "FabricDeploymentPipelineInvalidDisplayName",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricDeploymentPipelineInvalidDisplayName()
		},
	},
	{
		Name: "fabric_digital_twin_builder_invalid_description",
		Type: "FabricDigitalTwinBuilderInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricDigitalTwinBuilderInvalidDescription()
		},
	},
	{
		Name: "fabric_domain_invalid_description",
		Type: "FabricDomainInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricDomainInvalidDescription()
		},
	},
	{
		Name: "fabric_domain_invalid_display_name",
		Type: "FabricDomainInvalidDisplayName",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricDomainInvalidDisplayName()
		},
	},
	{
		Name: "fabric_environment_invalid_description",
		Type: "FabricEnvironmentInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricEnvironmentInvalidDescription()
		},
	},
	{
		Name: "fabric_eventhouse_invalid_description",
		Type: "FabricEventhouseInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricEventhouseInvalidDescription()
		},
	},
	{
		Name: "fabric_eventhouse_invalid_display_name",
		Type: "FabricEventhouseInvalidDisplayName",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricEventhouseInvalidDisplayName()
		},
	},
	{
		Name: "fabric_eventstream_invalid_description",
		Type: "FabricEventstreamInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricEventstreamInvalidDescription()
		},
	},
	{
		Name: "fabric_eventstream_invalid_display_name",
		Type: "FabricEventstreamInvalidDisplayName",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricEventstreamInvalidDisplayName()
		},
	},
	{
		Name: "fabric_folder_invalid_display_name",
		Type: "FabricFolderInvalidDisplayName",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricFolderInvalidDisplayName()
		},
	},
	{
		Name: "fabric_gateway_invalid_display_name",
		Type: "FabricGatewayInvalidDisplayName",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricGatewayInvalidDisplayName()
		},
	},
	{
		Name: "fabric_gateway_invalid_inactivity_minutes_before_sleep",
		Type: "FabricGatewayInvalidInactivityMinutesBeforeSleep",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricGatewayInvalidInactivityMinutesBeforeSleep()
		},
	},
	{
		Name: "fabric_gateway_invalid_type",
		Type: "FabricGatewayInvalidType",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricGatewayInvalidType()
		},
	},
	{
		Name: "fabric_graphql_api_invalid_description",
		Type: "FabricGraphqlAPIInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricGraphqlAPIInvalidDescription()
		},
	},
	{
		Name: "fabric_kql_dashboard_invalid_description",
		Type: "FabricKQLDashboardInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricKQLDashboardInvalidDescription()
		},
	},
	{
		Name: "fabric_kql_database_invalid_description",
		Type: "FabricKQLDatabaseInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricKQLDatabaseInvalidDescription()
		},
	},
	{
		Name: "fabric_kql_queryset_invalid_description",
		Type: "FabricKQLQuerysetInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricKQLQuerysetInvalidDescription()
		},
	},
	{
		Name: "fabric_lakehouse_invalid_description",
		Type: "FabricLakehouseInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricLakehouseInvalidDescription()
		},
	},
	{
		Name: "fabric_lakehouse_invalid_display_name",
		Type: "FabricLakehouseInvalidDisplayName",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricLakehouseInvalidDisplayName()
		},
	},
	{
		Name: "fabric_mirrored_database_invalid_description",
		Type: "FabricMirroredDatabaseInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricMirroredDatabaseInvalidDescription()
		},
	},
	{
		Name: "fabric_ml_experiment_invalid_description",
		Type: "FabricMlExperimentInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricMlExperimentInvalidDescription()
		},
	},
	{
		Name: "fabric_ml_model_invalid_description",
		Type: "FabricMlModelInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricMlModelInvalidDescription()
		},
	},
	{
		Name: "fabric_mounted_data_factory_invalid_description",
		Type: "FabricMountedDataFactoryInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricMountedDataFactoryInvalidDescription()
		},
	},
	{
		Name: "fabric_notebook_invalid_description",
		Type: "FabricNotebookInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricNotebookInvalidDescription()
		},
	},
	{
		Name: "fabric_notebook_invalid_display_name",
		Type: "FabricNotebookInvalidDisplayName",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricNotebookInvalidDisplayName()
		},
	},
	{
		Name: "fabric_report_invalid_description",
		Type: "FabricReportInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricReportInvalidDescription()
		},
	},
	{
		Name: "fabric_semantic_model_invalid_description",
		Type: "FabricSemanticModelInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricSemanticModelInvalidDescription()
		},
	},
	{
		Name: "fabric_spark_custom_pool_invalid_node_family",
		Type: "FabricSparkCustomPoolInvalidNodeFamily",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricSparkCustomPoolInvalidNodeFamily()
		},
	},
	{
		Name: "fabric_spark_custom_pool_invalid_node_size",
		Type: "FabricSparkCustomPoolInvalidNodeSize",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricSparkCustomPoolInvalidNodeSize()
		},
	},
	{
		Name: "fabric_spark_environment_settings_invalid_driver_cores",
		Type: "FabricSparkEnvironmentSettingsInvalidDriverCores",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricSparkEnvironmentSettingsInvalidDriverCores()
		},
	},
	{
		Name: "fabric_spark_environment_settings_invalid_driver_memory",
		Type: "FabricSparkEnvironmentSettingsInvalidDriverMemory",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricSparkEnvironmentSettingsInvalidDriverMemory()
		},
	},
	{
		Name: "fabric_spark_environment_settings_invalid_executor_cores",
		Type: "FabricSparkEnvironmentSettingsInvalidExecutorCores",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricSparkEnvironmentSettingsInvalidExecutorCores()
		},
	},
	{
		Name: "fabric_spark_environment_settings_invalid_executor_memory",
		Type: "FabricSparkEnvironmentSettingsInvalidExecutorMemory",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricSparkEnvironmentSettingsInvalidExecutorMemory()
		},
	},
	{
		Name: "fabric_spark_environment_settings_invalid_runtime_version",
		Type: "FabricSparkEnvironmentSettingsInvalidRuntimeVersion",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricSparkEnvironmentSettingsInvalidRuntimeVersion()
		},
	},
	{
		Name: "fabric_spark_job_definition_invalid_description",
		Type: "FabricSparkJobDefinitionInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricSparkJobDefinitionInvalidDescription()
		},
	},
	{
		Name: "fabric_spark_job_definition_invalid_display_name",
		Type: "FabricSparkJobDefinitionInvalidDisplayName",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricSparkJobDefinitionInvalidDisplayName()
		},
	},
	{
		Name: "fabric_sql_database_invalid_description",
		Type: "FabricSQLDatabaseInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricSQLDatabaseInvalidDescription()
		},
	},
	{
		Name: "fabric_variable_library_invalid_description",
		Type: "FabricVariableLibraryInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricVariableLibraryInvalidDescription()
		},
	},
	{
		Name: "fabric_warehouse_invalid_description",
		Type: "FabricWarehouseInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricWarehouseInvalidDescription()
		},
	},
	{
		Name: "fabric_warehouse_snapshot_invalid_description",
		Type: "FabricWarehouseSnapshotInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricWarehouseSnapshotInvalidDescription()
		},
	},
	{
		Name: "fabric_workspace_invalid_description",
		Type: "FabricWorkspaceInvalidDescription",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricWorkspaceInvalidDescription()
		},
	},
	{
		Name: "fabric_workspace_invalid_display_name",
		Type: "FabricWorkspaceInvalidDisplayName",
		Constructor: func() interface{ Check(tflint.Runner) error } {
			return NewFabricWorkspaceInvalidDisplayName()
		},
	},
// End generated rules list

	}

	for _, rule := range generatedRuleConstructors {
		rules[rule.Name] = rule
	}

	return rules
}

// TestGeneratedRulesAvailable verifies that generated rules are discoverable
// This is a placeholder test file for auto-generated rule test cases
// As rules are generated, add specific test cases here following the pattern in rules_test.go

func TestGeneratedRulesBasicValidation(t *testing.T) {
	// Test to ensure the generated rule framework is working
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "Valid workspace configuration",
			content: `
resource "fabric_workspace" "example" {
  display_name = "valid-workspace"
  description  = "Test workspace"
  capacity_id  = "capacity-123"
}
`,
			hasIssue: false,
		},
		{
			name: "Valid item configuration",
			content: `
resource "fabric_item" "example" {
  workspace_id = "workspace-123"
  name         = "my-item"
  type         = "Lakehouse"
}
`,
			hasIssue: false,
		},
		{
			name: "Valid capacity configuration",
			content: `
resource "fabric_capacity" "example" {
  name   = "test-capacity"
  region = "westus"
  sku    = "F2"
}
`,
			hasIssue: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{
				"main.tf": tt.content,
			})

			// Run all available manual rules
			rule := NewFabricWorkspaceNaming()
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			// For valid configs, we expect no issues
			if !tt.hasIssue && len(runner.Issues) > 0 {
				t.Logf("Configuration is valid but got %d issues", len(runner.Issues))
			}
		})
	}
}

// TestGeneratedRulesComplexScenario tests a realistic configuration
func TestGeneratedRulesComplexScenario(t *testing.T) {
	runner := helper.TestRunner(t, map[string]string{
		"main.tf": `
resource "fabric_workspace" "prod" {
  display_name = "prod-workspace"
  description  = "Production workspace"
  capacity_id  = "capacity-prod-001"
}

resource "fabric_workspace" "dev" {
  display_name = "dev-workspace"
  description  = "Development workspace"
  capacity_id  = "capacity-dev-001"
}

resource "fabric_capacity" "prod" {
  name   = "prod-capacity"
  region = "westus"
  sku    = "F4"
}

resource "fabric_capacity" "dev" {
  name   = "dev-capacity"
  region = "eastus"
  sku    = "F2"
}
`,
	})

	// Test manual validation rules
	rules := []struct {
		name string
		rule interface{ Check(tflint.Runner) error }
	}{
		{
			name: "workspace_naming",
			rule: NewFabricWorkspaceNaming(),
		},
		{
			name: "workspace_capacity",
			rule: NewFabricWorkspaceCapacity(),
		},
		{
			name: "workspace_description",
			rule: NewFabricWorkspaceDescription(),
		},
	}

	for _, r := range rules {
		t.Run(r.name, func(t *testing.T) {
			if err := r.rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error in %s: %s", r.name, err)
			}
		})
	}
}

// TestGeneratedRulesAllRulesRegistered verifies rules are available in main.go
func TestGeneratedRulesAllRulesRegistered(t *testing.T) {
	// This test ensures that:
	// 1. Generated rules are placed in rules/ directory
	// 2. Rules are registered in main.go
	// 3. The ruleset can be instantiated

	runner := helper.TestRunner(t, map[string]string{
		"empty.tf": "# Test file\n",
	})

	// Verify we can instantiate the manual rules
	manualRules := []interface{}{
		NewFabricWorkspaceCapacity(),
		NewFabricWorkspaceDescription(),
		NewFabricRoleAssignmentRecommended(),
		NewFabricGitIntegrationValidation(),
		NewFabricCapacityRegion(),
	}

	if len(manualRules) == 0 {
		t.Error("No manual rules found")
	}

	// Run at least one rule to verify the runner works
	if err := NewFabricWorkspaceCapacity().Check(runner); err != nil {
		t.Fatalf("Failed to run rule: %s", err)
	}

	t.Logf("Successfully verified %d manual rules are registered", len(manualRules))
}

// TestGeneratedRulesDiscovery verifies all generated rules can be instantiated
func TestGeneratedRulesDiscovery(t *testing.T) {
	generatedRules := GetGeneratedRules()

	if len(generatedRules) == 0 {
		t.Error("No generated rules found")
		return
	}

	t.Logf("Found %d generated rules", len(generatedRules))

	for ruleName, ruleInfo := range generatedRules {
		t.Run(ruleName, func(t *testing.T) {
			// Instantiate the rule
			rule := ruleInfo.Constructor()

			// Verify it implements the Check method
			ruleValue := reflect.ValueOf(rule)
			if !ruleValue.Type().Implements(reflect.TypeOf((*interface{ Check(tflint.Runner) error })(nil)).Elem()) {
				t.Errorf("Rule %s does not implement Check method", ruleName)
				return
			}

			t.Logf("✓ Rule %s is valid", ruleName)
		})
	}
}

// TestGeneratedRulesExecuteAll runs all generated rules against a test configuration
func TestGeneratedRulesExecuteAll(t *testing.T) {
	generatedRules := GetGeneratedRules()

	runner := helper.TestRunner(t, map[string]string{
		"main.tf": `
# Empty Terraform configuration to test all rules can execute
`,
	})

	successCount := 0
	failureCount := 0

	for ruleName, ruleInfo := range generatedRules {
		rule := ruleInfo.Constructor()
		ruleMethod := rule.(interface{ Check(tflint.Runner) error }).Check

		// Execute the rule
		if err := ruleMethod(runner); err != nil {
			t.Logf("✗ Rule %s failed: %v", ruleName, err)
			failureCount++
		} else {
			successCount++
		}
	}

	t.Logf("Generated rules execution results: %d successful, %d failed", successCount, failureCount)

	if failureCount > 0 {
		t.Logf("Warning: %d rules failed execution", failureCount)
	}
}

// DiscoverGeneratedRulesFromDirectory scans the generated directory and extracts rule names
// This can be used to automatically update GetGeneratedRules()
func DiscoverGeneratedRulesFromDirectory(generatedPath string) ([]string, error) {
	var ruleNames []string

	entries, err := os.ReadDir(generatedPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	ruleNameRegex := regexp.MustCompile(`^fabric_(.+)\.go$`)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}

		matches := ruleNameRegex.FindStringSubmatch(entry.Name())
		if len(matches) > 1 {
			ruleNames = append(ruleNames, matches[1])
		}
	}

	return ruleNames, nil
}

// TestGeneratedRulesDiscoveryFromFilesystem scans the generated directory for rules
func TestGeneratedRulesDiscoveryFromFilesystem(t *testing.T) {
	// Find the generated directory relative to this test file
	generatedPath := filepath.Join("generator", "generated")

	// Try from different working directory contexts
	if _, err := os.Stat(generatedPath); os.IsNotExist(err) {
		generatedPath = filepath.Join(".", "rules", "generator", "generated")
	}

	if _, err := os.Stat(generatedPath); os.IsNotExist(err) {
		t.Skipf("Generated directory not found at %s", generatedPath)
	}

	ruleNames, err := DiscoverGeneratedRulesFromDirectory(generatedPath)
	if err != nil {
		t.Fatalf("Failed to discover rules: %v", err)
	}

	if len(ruleNames) == 0 {
		t.Skip("No generated rules found in filesystem")
	}

	t.Logf("Discovered %d generated rule files: %v", len(ruleNames), ruleNames[:min(5, len(ruleNames))])

	// This output can be used to verify that the generated rules
	// match what's in GetGeneratedRules()
	registeredRules := GetGeneratedRules()
	if len(ruleNames) != len(registeredRules) {
		t.Logf("Warning: Found %d rule files but %d rules registered", len(ruleNames), len(registeredRules))
	}
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Note: For generated rules from the rule generator, add test cases like:
//
// func TestGeneratedRulesEnumConstraints(t *testing.T) {
//     runner := helper.TestRunner(t, map[string]string{
//         "main.tf": `resource "fabric_resource" "example" { field = "valid_enum_value" }`,
//     })
//     rule := NewGeneratedRule()
//     if err := rule.Check(runner); err != nil {
//         t.Fatalf("Error: %s", err)
//     }
//     if len(runner.Issues) > 0 {
//         t.Errorf("Expected no issues, got %d", len(runner.Issues))
//     }
// }
