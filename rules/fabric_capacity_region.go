package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hcl"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
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
	return false // Disabled by default, enable if needed
}

func (r *FabricCapacityRegion) Severity() string {
	return tflint.WARNING
}

func (r *FabricCapacityRegion) Link() string {
	return "https://learn.microsoft.com/en-us/fabric/admin/capacity-settings"
}

func (r *FabricCapacityRegion) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourcesByType("fabric_capacity")
	if err != nil {
		return err
	}

	validRegions := map[string]bool{
		"eastus":      true,
		"westus":      true,
		"westeurope":  true,
		"eastasia":    true,
		"southeastasia": true,
		"uksouth":     true,
		"australiaeast": true,
		"canadacentral": true,
		"brazilsouth": true,
	}

	for _, resource := range resources {
		var region string
		err := resource.GetAttribute("region", &region)
		if err == nil && region != "" {
			if !validRegions[region] {
				runner.EmitIssue(
					r,
					fmt.Sprintf("Region %s may not be available. Please verify with Microsoft Fabric documentation", region),
					resource.GetNameRange(),
				)
			}
		}
	}

	return nil
}