module github.com/RuneORakeie/tflint-ruleset-fabric/tools/apispec-rule-gen

go 1.25

require (
	// Needed to build the plugin (generated rules, provider entry)
	github.com/terraform-linters/tflint-plugin-sdk v0.23.1

	// Needed by your rule-generator under tools/apispec-rule-gen
	github.com/hashicorp/hcl/v2 v2.19.1
	github.com/zclconf/go-cty v1.14.1
)

require (
	// Indirects used by the rule-generator
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	golang.org/x/text v0.14.0 // indirect
)
