package apispec

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GeneratedRuleInfo contains metadata about a generated rule
type GeneratedRuleInfo struct {
	Name        string
	Type        string
	Constructor func() interface{ Check(tflint.Runner) error }
}

// GetGeneratedRules returns a map of all generated rule constructors
func GetGeneratedRules() map[string]GeneratedRuleInfo {
	rules := make(map[string]GeneratedRuleInfo)

	// All 54 generated rules in apispec package
	generatedRuleConstructors := []GeneratedRuleInfo{
		{
			Name:        "fabric_activator_invalid_description",
			Type:        "FabricActivatorInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricActivatorInvalidDescription() },
		},
		{
			Name:        "fabric_apache_airflow_job_invalid_description",
			Type:        "FabricApacheAirflowJobInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricApacheAirflowJobInvalidDescription() },
		},
		{
			Name:        "fabric_connection_invalid_connectivity_type",
			Type:        "FabricConnectionInvalidConnectivityType",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricConnectionInvalidConnectivityType() },
		},
		{
			Name:        "fabric_connection_invalid_display_name",
			Type:        "FabricConnectionInvalidDisplayName",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricConnectionInvalidDisplayName() },
		},
		{
			Name:        "fabric_connection_invalid_privacy_level",
			Type:        "FabricConnectionInvalidPrivacyLevel",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricConnectionInvalidPrivacyLevel() },
		},
		{
			Name:        "fabric_copy_job_invalid_description",
			Type:        "FabricCopyJobInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricCopyJobInvalidDescription() },
		},
		{
			Name:        "fabric_copy_job_invalid_display_name",
			Type:        "FabricCopyJobInvalidDisplayName",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricCopyJobInvalidDisplayName() },
		},
		{
			Name:        "fabric_data_pipeline_invalid_description",
			Type:        "FabricDataPipelineInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricDataPipelineInvalidDescription() },
		},
		{
			Name:        "fabric_data_pipeline_invalid_display_name",
			Type:        "FabricDataPipelineInvalidDisplayName",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricDataPipelineInvalidDisplayName() },
		},
		{
			Name:        "fabric_dataflow_invalid_description",
			Type:        "FabricDataflowInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricDataflowInvalidDescription() },
		},
		{
			Name:        "fabric_dataflow_invalid_display_name",
			Type:        "FabricDataflowInvalidDisplayName",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricDataflowInvalidDisplayName() },
		},
		{
			Name:        "fabric_deployment_pipeline_invalid_description",
			Type:        "FabricDeploymentPipelineInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricDeploymentPipelineInvalidDescription() },
		},
		{
			Name:        "fabric_deployment_pipeline_invalid_display_name",
			Type:        "FabricDeploymentPipelineInvalidDisplayName",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricDeploymentPipelineInvalidDisplayName() },
		},
		{
			Name:        "fabric_digital_twin_builder_invalid_description",
			Type:        "FabricDigitalTwinBuilderInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricDigitalTwinBuilderInvalidDescription() },
		},
		{
			Name:        "fabric_domain_invalid_description",
			Type:        "FabricDomainInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricDomainInvalidDescription() },
		},
		{
			Name:        "fabric_domain_invalid_display_name",
			Type:        "FabricDomainInvalidDisplayName",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricDomainInvalidDisplayName() },
		},
		{
			Name:        "fabric_environment_invalid_description",
			Type:        "FabricEnvironmentInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricEnvironmentInvalidDescription() },
		},
		{
			Name:        "fabric_eventhouse_invalid_description",
			Type:        "FabricEventhouseInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricEventhouseInvalidDescription() },
		},
		{
			Name:        "fabric_eventhouse_invalid_display_name",
			Type:        "FabricEventhouseInvalidDisplayName",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricEventhouseInvalidDisplayName() },
		},
		{
			Name:        "fabric_eventstream_invalid_description",
			Type:        "FabricEventstreamInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricEventstreamInvalidDescription() },
		},
		{
			Name:        "fabric_eventstream_invalid_display_name",
			Type:        "FabricEventstreamInvalidDisplayName",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricEventstreamInvalidDisplayName() },
		},
		{
			Name:        "fabric_folder_invalid_display_name",
			Type:        "FabricFolderInvalidDisplayName",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricFolderInvalidDisplayName() },
		},
		{
			Name:        "fabric_gateway_invalid_display_name",
			Type:        "FabricGatewayInvalidDisplayName",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricGatewayInvalidDisplayName() },
		},
		{
			Name: "fabric_gateway_invalid_inactivity_minutes_before_sleep",
			Type: "FabricGatewayInvalidInactivityMinutesBeforeSleep",
			Constructor: func() interface{ Check(tflint.Runner) error } {
				return NewFabricGatewayInvalidInactivityMinutesBeforeSleep()
			},
		},
		{
			Name:        "fabric_gateway_invalid_type",
			Type:        "FabricGatewayInvalidType",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricGatewayInvalidType() },
		},
		{
			Name:        "fabric_graphql_api_invalid_description",
			Type:        "FabricGraphqlAPIInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricGraphqlAPIInvalidDescription() },
		},
		{
			Name:        "fabric_kql_dashboard_invalid_description",
			Type:        "FabricKQLDashboardInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricKQLDashboardInvalidDescription() },
		},
		{
			Name:        "fabric_kql_database_invalid_description",
			Type:        "FabricKQLDatabaseInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricKQLDatabaseInvalidDescription() },
		},
		{
			Name:        "fabric_kql_queryset_invalid_description",
			Type:        "FabricKQLQuerysetInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricKQLQuerysetInvalidDescription() },
		},
		{
			Name:        "fabric_lakehouse_invalid_description",
			Type:        "FabricLakehouseInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricLakehouseInvalidDescription() },
		},
		{
			Name:        "fabric_lakehouse_invalid_display_name",
			Type:        "FabricLakehouseInvalidDisplayName",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricLakehouseInvalidDisplayName() },
		},
		{
			Name:        "fabric_mirrored_database_invalid_description",
			Type:        "FabricMirroredDatabaseInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricMirroredDatabaseInvalidDescription() },
		},
		{
			Name:        "fabric_ml_experiment_invalid_description",
			Type:        "FabricMlExperimentInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricMlExperimentInvalidDescription() },
		},
		{
			Name:        "fabric_ml_model_invalid_description",
			Type:        "FabricMlModelInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricMlModelInvalidDescription() },
		},
		{
			Name:        "fabric_mounted_data_factory_invalid_description",
			Type:        "FabricMountedDataFactoryInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricMountedDataFactoryInvalidDescription() },
		},
		{
			Name:        "fabric_notebook_invalid_description",
			Type:        "FabricNotebookInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricNotebookInvalidDescription() },
		},
		{
			Name:        "fabric_notebook_invalid_display_name",
			Type:        "FabricNotebookInvalidDisplayName",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricNotebookInvalidDisplayName() },
		},
		{
			Name:        "fabric_report_invalid_description",
			Type:        "FabricReportInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricReportInvalidDescription() },
		},
		{
			Name:        "fabric_semantic_model_invalid_description",
			Type:        "FabricSemanticModelInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricSemanticModelInvalidDescription() },
		},
		{
			Name:        "fabric_spark_custom_pool_invalid_node_family",
			Type:        "FabricSparkCustomPoolInvalidNodeFamily",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricSparkCustomPoolInvalidNodeFamily() },
		},
		{
			Name:        "fabric_spark_custom_pool_invalid_node_size",
			Type:        "FabricSparkCustomPoolInvalidNodeSize",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricSparkCustomPoolInvalidNodeSize() },
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
			Name:        "fabric_spark_job_definition_invalid_description",
			Type:        "FabricSparkJobDefinitionInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricSparkJobDefinitionInvalidDescription() },
		},
		{
			Name:        "fabric_spark_job_definition_invalid_display_name",
			Type:        "FabricSparkJobDefinitionInvalidDisplayName",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricSparkJobDefinitionInvalidDisplayName() },
		},
		{
			Name:        "fabric_sql_database_invalid_description",
			Type:        "FabricSQLDatabaseInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricSQLDatabaseInvalidDescription() },
		},
		{
			Name:        "fabric_variable_library_invalid_description",
			Type:        "FabricVariableLibraryInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricVariableLibraryInvalidDescription() },
		},
		{
			Name:        "fabric_warehouse_invalid_description",
			Type:        "FabricWarehouseInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricWarehouseInvalidDescription() },
		},
		{
			Name:        "fabric_warehouse_snapshot_invalid_description",
			Type:        "FabricWarehouseSnapshotInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricWarehouseSnapshotInvalidDescription() },
		},
		{
			Name:        "fabric_workspace_invalid_description",
			Type:        "FabricWorkspaceInvalidDescription",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricWorkspaceInvalidDescription() },
		},
		{
			Name:        "fabric_workspace_invalid_display_name",
			Type:        "FabricWorkspaceInvalidDisplayName",
			Constructor: func() interface{ Check(tflint.Runner) error } { return NewFabricWorkspaceInvalidDisplayName() },
		},
	}

	for _, rule := range generatedRuleConstructors {
		rules[rule.Name] = rule
	}

	return rules
}

// TestGeneratedRulesBasicConfiguration tests that generated rules don't error on empty/valid configs
func TestGeneratedRulesBasicConfiguration(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		hasIssue bool
	}{
		{
			name: "Empty configuration",
			content: `
# Empty Terraform configuration
`,
			hasIssue: false,
		},
		{
			name: "Valid workspace configuration",
			content: `
resource "fabric_workspace" "example" {
  display_name = "test-workspace"
  description  = "Test workspace"
  capacity_id  = "test-capacity-id"
}
`,
			hasIssue: false,
		},
	}

	generatedRules := GetGeneratedRules()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{
				"main.tf": tt.content,
			})

			// Run one sample rule to verify basic functionality
			if rule, exists := generatedRules["fabric_workspace_invalid_description"]; exists {
				if err := rule.Constructor().Check(runner); err != nil {
					t.Fatalf("Unexpected error: %s", err)
				}
			}
		})
	}
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
		ruleMethod := rule.Check

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
		t.Errorf("%d rules failed execution", failureCount)
	}
}

// DiscoverGeneratedRulesFromDirectory scans the apispec directory and extracts rule names
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

		// Skip test files and provider.go
		if strings.HasSuffix(entry.Name(), "_test.go") || entry.Name() == "provider.go" {
			continue
		}

		matches := ruleNameRegex.FindStringSubmatch(entry.Name())
		if len(matches) > 1 {
			ruleName := "fabric_" + matches[1]
			ruleNames = append(ruleNames, ruleName)
		}
	}

	return ruleNames, nil
}

// TestGeneratedRulesDiscoveryFromFilesystem scans the apispec directory for rules
func TestGeneratedRulesDiscoveryFromFilesystem(t *testing.T) {
	// This test runs from rules/apispec/test_generated_rules.go
	// So current directory is the apispec directory
	generatedPath := "."

	ruleNames, err := DiscoverGeneratedRulesFromDirectory(generatedPath)
	if err != nil {
		t.Fatalf("Failed to discover rules: %v", err)
	}

	if len(ruleNames) == 0 {
		t.Skip("No generated rules found in filesystem")
	}

	t.Logf("Discovered %d generated rule files", len(ruleNames))

	// Verify that registered rules match what's in filesystem
	registeredRules := GetGeneratedRules()
	if len(ruleNames) != len(registeredRules) {
		t.Errorf("Found %d rule files but %d rules registered", len(ruleNames), len(registeredRules))

		// Show which rules are missing
		registered := make(map[string]bool)
		for name := range registeredRules {
			registered[name] = true
		}

		for _, name := range ruleNames {
			if !registered[name] {
				t.Logf("  Missing in registry: %s", name)
			}
		}
	} else {
		t.Logf("✓ All %d filesystem rules are registered", len(ruleNames))
	}
}
