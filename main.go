// main.go - Plugin entry point
package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"

	"github.com/RuneORakeie/tflint-ruleset-fabric/fabric"
)

var (
	Sha1ver = ""
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: fabric.NewRuleSet(),
	})
}
