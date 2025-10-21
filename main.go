// main.go - Plugin entry point
package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/RuneORakeie/tflint-ruleset-fabric/rules"
)

var (
	Version = "0.1.0"
	Sha1ver = ""
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "fabric",
			Version: Version,
			Rules: []tflint.Rule{
				rules.NewFabricWorkspaceNaming(),
				rules.NewFabricWorkspaceCapacity(),
				rules.NewFabricWorkspaceDescription(),
				rules.NewFabricRoleAssignmentPrincipals(),
				rules.NewFabricGitIntegrationValidation(),
				rules.NewFabricCapacityRegion(),
			},
		},
	})
}