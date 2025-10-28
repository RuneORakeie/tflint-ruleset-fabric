package rules

import (
	"fmt"
	"strings"

	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// FabricWorkspaceGitInitializationStrategy validates initialization_strategy values
type FabricWorkspaceGitInitializationStrategy struct {
	tflint.DefaultRule
}

func NewFabricWorkspaceGitInitializationStrategy() *FabricWorkspaceGitInitializationStrategy {
	return &FabricWorkspaceGitInitializationStrategy{}
}

func (r *FabricWorkspaceGitInitializationStrategy) Name() string {
	return "fabric_workspace_git_initialization_strategy_valid"
}

func (r *FabricWorkspaceGitInitializationStrategy) Enabled() bool {
	return true
}

func (r *FabricWorkspaceGitInitializationStrategy) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *FabricWorkspaceGitInitializationStrategy) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *FabricWorkspaceGitInitializationStrategy) Check(runner tflint.Runner) error {
	resourceContent, err := runner.GetResourceContent("fabric_workspace_git", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "initialization_strategy"},
		},
	}, nil)
	if err != nil {
		return err
	}

	validStrategies := map[string]bool{
		"PreferRemote":    true,
		"PreferWorkspace": true,
	}

	for _, resource := range resourceContent.Blocks {
		if attr, exists := resource.Body.Attributes["initialization_strategy"]; exists && attr.Expr != nil {
			var strategy string
			if err := runner.EvaluateExpr(attr.Expr, &strategy, nil); err == nil && strategy != "" {
				if !validStrategies[strategy] {
					validStrategiesList := []string{"PreferRemote", "PreferWorkspace"}
					runner.EmitIssue(
						r,
						fmt.Sprintf("Invalid initialization_strategy '%s'. Must be one of: %s", strategy, strings.Join(validStrategiesList, ", ")),
						attr.Range,
					)
				}
			}
		}
	}

	return nil
}
