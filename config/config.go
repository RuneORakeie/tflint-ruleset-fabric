package config

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
)

// Config represents the plugin configuration
type Config struct {
	Preset           string            `hcl:"preset,optional"`
	NamingConvention *NamingConvention `hcl:"naming_convention,block"`
}

// NamingConvention defines naming patterns for resources
type NamingConvention struct {
	// General settings that apply to all resources
	Pattern string `hcl:"pattern,optional"` // General regex pattern for all display_name attributes
	Format  string `hcl:"format,optional"`  // snake_case, kebab-case, PascalCase, camelCase
	Prefix  string `hcl:"prefix,optional"`  // General prefix for all resources
	Suffix  string `hcl:"suffix,optional"`  // General suffix for all resources

	// Resource-specific overrides using resource type names
	PatternOverrides map[string]string `hcl:"pattern_overrides,optional"`
	PrefixOverrides  map[string]string `hcl:"prefix_overrides,optional"`
	SuffixOverrides  map[string]string `hcl:"suffix_overrides,optional"`
}

// Decode parses the configuration from HCL
func (c *Config) Decode(body *hclext.BodyContent) error {
	// Parse preset attribute
	if attr, exists := body.Attributes["preset"]; exists {
		var preset string
		if err := gohcl.DecodeExpression(attr.Expr, nil, &preset); err != nil {
			return err
		}
		c.Preset = preset
	}

	// Parse naming_convention block
	for _, block := range body.Blocks {
		if block.Type == "naming_convention" {
			nc := &NamingConvention{
				PatternOverrides: make(map[string]string),
				PrefixOverrides:  make(map[string]string),
				SuffixOverrides:  make(map[string]string),
			}

			// Parse general settings
			if attr, exists := block.Body.Attributes["pattern"]; exists {
				if err := gohcl.DecodeExpression(attr.Expr, nil, &nc.Pattern); err != nil {
					return err
				}
			}
			if attr, exists := block.Body.Attributes["format"]; exists {
				if err := gohcl.DecodeExpression(attr.Expr, nil, &nc.Format); err != nil {
					return err
				}
			}
			if attr, exists := block.Body.Attributes["prefix"]; exists {
				if err := gohcl.DecodeExpression(attr.Expr, nil, &nc.Prefix); err != nil {
					return err
				}
			}
			if attr, exists := block.Body.Attributes["suffix"]; exists {
				if err := gohcl.DecodeExpression(attr.Expr, nil, &nc.Suffix); err != nil {
					return err
				}
			}

			// Parse override maps
			if attr, exists := block.Body.Attributes["pattern_overrides"]; exists {
				if err := gohcl.DecodeExpression(attr.Expr, nil, &nc.PatternOverrides); err != nil {
					return err
				}
			}
			if attr, exists := block.Body.Attributes["prefix_overrides"]; exists {
				if err := gohcl.DecodeExpression(attr.Expr, nil, &nc.PrefixOverrides); err != nil {
					return err
				}
			}
			if attr, exists := block.Body.Attributes["suffix_overrides"]; exists {
				if err := gohcl.DecodeExpression(attr.Expr, nil, &nc.SuffixOverrides); err != nil {
					return err
				}
			}

			c.NamingConvention = nc
		}
	}

	// Validate preset value
	if c.Preset != "" {
		validPresets := map[string]bool{
			"minimal":     true,
			"recommended": true,
			"all":         true,
		}
		if !validPresets[c.Preset] {
			return fmt.Errorf("invalid preset %q: must be one of minimal, recommended, all", c.Preset)
		}
	} else {
		// Default to "all" if not specified
		c.Preset = "all"
	}

	// Validate naming convention format
	if c.NamingConvention != nil && c.NamingConvention.Format != "" {
		validFormats := map[string]bool{
			"snake_case": true,
			"kebab-case": true,
			"PascalCase": true,
			"camelCase":  true,
		}
		if !validFormats[c.NamingConvention.Format] {
			return fmt.Errorf("invalid format %q: must be one of snake_case, kebab-case, PascalCase, camelCase", c.NamingConvention.Format)
		}
	}

	return nil
}

// GetPreset returns the configured preset, defaulting to "all"
func (c *Config) GetPreset() string {
	if c.Preset == "" {
		return "all"
	}
	return c.Preset
}

// GetPatternForResource returns the naming pattern for a specific resource type
// Falls back to general pattern if resource-specific pattern is not set
func (nc *NamingConvention) GetPatternForResource(resourceType string) string {
	if nc == nil {
		return ""
	}

	// Check for resource-specific override
	if pattern, exists := nc.PatternOverrides[resourceType]; exists {
		return pattern
	}

	// Fall back to general pattern
	return nc.Pattern
}

// GetPrefixForResource returns the naming prefix for a specific resource type
// Falls back to general prefix if resource-specific prefix is not set
func (nc *NamingConvention) GetPrefixForResource(resourceType string) string {
	if nc == nil {
		return ""
	}

	// Check for resource-specific override
	if prefix, exists := nc.PrefixOverrides[resourceType]; exists {
		return prefix
	}

	// Fall back to general prefix
	return nc.Prefix
}

// GetSuffixForResource returns the naming suffix for a specific resource type
// Falls back to general suffix if resource-specific suffix is not set
func (nc *NamingConvention) GetSuffixForResource(resourceType string) string {
	if nc == nil {
		return ""
	}

	// Check for resource-specific override
	if suffix, exists := nc.SuffixOverrides[resourceType]; exists {
		return suffix
	}

	// Fall back to general suffix
	return nc.Suffix
}
