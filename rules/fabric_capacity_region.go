package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
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

func (r *FabricCapacityRegion) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *FabricCapacityRegion) Link() string {
	return "https://learn.microsoft.com/en-us/fabric/admin/capacity-settings"
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

	validRegions := map[string]bool{
		"eastus":        true,
		"westus":        true,
		"westeurope":    true,
		"eastasia":      true,
		"southeastasia": true,
		"uksouth":       true,
		"australiaeast": true,
		"canadacentral": true,
		"brazilsouth":   true,
	}

	if attr, exists := resources.Attributes["region"]; exists && attr.Expr != nil {
		var region string
		if err := runner.EvaluateExpr(attr.Expr, &region, nil); err == nil && region != "" {
			if !validRegions[region] {
				runner.EmitIssue(
					r,
					fmt.Sprintf("Region %s may not be available. Please verify with Microsoft Fabric documentation", region),
					attr.Range,
				)
			}
		}
	}

	return nil
}
