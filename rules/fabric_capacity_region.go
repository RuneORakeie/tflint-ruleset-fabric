package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricCapacityRegion validates capacity region
type FabricCapacityRegion struct {
	tflint.DefaultRule
}

func NewFabricCapacityRegion() *FabricCapacityRegion {
	return &FabricCapacityRegion{}
}

func (r *FabricCapacityRegion) Name() string {
	return "fabric_capacity_region_valid"
}

func (r *FabricCapacityRegion) Enabled() bool {
	return true // Disabled by default, enable if needed
}

func (r *FabricCapacityRegion) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *FabricCapacityRegion) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *FabricCapacityRegion) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("fabric_capacity", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "region"},
		},
	}, nil)
	if err != nil {
		return err
	}

	// Regions where all Fabric workloads are available (as of Sept 2025)
	// Source: https://learn.microsoft.com/en-us/fabric/admin/region-availability
	allWorkloadsRegions := map[string]bool{
		// Americas
		"brazilsouth":    true,
		"canadacentral":  true,
		"canadaeast":     true,
		"centralus":      true,
		"eastus":         true,
		"eastus2":        true,
		"mexicocentral":  true,
		"northcentralus": true,
		"southcentralus": true,
		"westus":         true,
		"westus2":        true,
		"westus3":        true,
		// Europe
		"northeurope":        true,
		"westeurope":         true,
		"francecentral":      true,
		"germanywestcentral": true,
		"italynorth":         true,
		"norwayeast":         true,
		"polandcentral":      true,
		"spaincentral":       true,
		"swedencentral":      true,
		"switzerlandnorth":   true,
		"switzerlandwest":    true,
		"uksouth":            true,
		"ukwest":             true,
		// Middle East & Africa
		"uaenorth":         true,
		"southafricanorth": true,
		// Asia Pacific
		"australiaeast":      true,
		"australiasoutheast": true,
		"centralindia":       true,
		"eastasia":           true,
		"israelcentral":      true,
		"japaneast":          true,
		"japanwest":          true,
		"southeastasia":      true,
		"southindia":         true,
		"koreacentral":       true,
	}

	for _, resource := range resources.Blocks {
		if attr, exists := resource.Body.Attributes["region"]; exists && attr.Expr != nil {
			var region string
			if err := runner.EvaluateExpr(attr.Expr, &region, nil); err == nil && region != "" {
				if !allWorkloadsRegions[region] {
					if err := runner.EmitIssue(
						r,
						fmt.Sprintf("Region '%s' may not support all Fabric workloads. Some features might be unavailable. Verify region availability at https://learn.microsoft.com/en-us/fabric/admin/region-availability", region),
						attr.Range,
					); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
