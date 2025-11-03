package fabric

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/RuneORakeie/tflint-ruleset-fabric/config"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
	"github.com/RuneORakeie/tflint-ruleset-fabric/rules"
)

// RuleSet is the custom ruleset for Fabric
type RuleSet struct {
	tflint.BuiltinRuleSet
	config *config.Config
}

// NewRuleSet creates a new RuleSet instance
func NewRuleSet() *RuleSet {
	return &RuleSet{
		BuiltinRuleSet: tflint.BuiltinRuleSet{
			Name:    "fabric",
			Version: project.Version,
			Rules:   rules.PresetRules["all"], // Default to all rules
		},
		config: &config.Config{
			Preset: "all", // Default preset
		},
	}
}

// ApplyConfig applies the parsed configuration
func (r *RuleSet) ApplyConfig(body *hclext.BodyContent) error {
	r.config = &config.Config{}
	if err := r.config.Decode(body); err != nil {
		return err
	}

	// Update rules based on preset
	preset := r.config.GetPreset()
	if rulesForPreset, ok := rules.PresetRules[preset]; ok {
		r.Rules = rulesForPreset
	}

	return nil
}

// GetConfig returns the current configuration
func (r *RuleSet) GetConfig() *config.Config {
	return r.config
}
