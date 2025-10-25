# fabric_capacity_region_valid

Validates that the capacity region is one where all Microsoft Fabric workloads are available.

## Example

```hcl
resource "fabric_capacity" "example" {
  display_name = "My Capacity"
  region       = "eastus" # Valid - all workloads available
  sku          = "F2"
}

resource "fabric_capacity" "limited" {
  display_name = "My Capacity"
  region       = "francesouth" # Warning - only Power BI available in this region
  sku          = "F2"
}
```

## Why

Microsoft Fabric workloads are not uniformly available across all Azure regions. While some regions support all Fabric capabilities, others may only support Power BI or have limited workload availability. Using a region with limited support may result in:

- Missing features or workloads (e.g., Synapse Data Engineering, Real-Time Intelligence)
- Inability to create certain resource types
- Unexpected deployment failures

This rule warns you when selecting a region that may not support all Fabric workloads, helping you make informed decisions about capacity placement.

## Validation Rules

**Regions where all workloads are supported:**

### Americas
- Brazil South
- Canada Central
- Canada East
- Central US
- East US
- East US 2
- Mexico Central
- North Central US
- South Central US
- West US
- West US 2
- West US 3

### Europe
- North Europe
- West Europe
- France Central
- Germany West Central
- Italy North
- Norway East
- Poland Central
- Spain Central
- Sweden Central
- Switzerland North
- Switzerland West
- UK South
- UK West

### Middle East & Africa
- UAE North
- South Africa North

### Asia Pacific
- Australia East
- Australia Southeast
- Central India
- East Asia
- Israel Central
- Japan East
- Japan West
- Southeast Asia
- South India
- Korea Central

**Note:** This list is based on the [official Microsoft Fabric region availability documentation](https://learn.microsoft.com/en-us/fabric/admin/region-availability) (as of September 2025). Some regions may have limitations for specific features even within the "all workloads" category (e.g., Fabric SQL database, Healthcare Solutions, User Data Functions).

## How to Fix

Choose a region from the "all workloads" list above, or verify the specific workload availability for your chosen region at the [Microsoft Fabric region availability page](https://learn.microsoft.com/en-us/fabric/admin/region-availability).

```hcl
resource "fabric_capacity" "production" {
  display_name = "Production Capacity"
  region       = "eastus"  # All workloads supported
  sku          = "F64"
}
```

## Configuration

```hcl
rule "fabric_capacity_region_valid" {
  enabled = true
}
```

## Attributes

| Name | Enabled | Severity | 
|------|---------|----------|
| fabric_capacity_region_valid | false | warning |

**Note:** This rule is disabled by default. Enable it if you want to receive warnings about potential workload availability limitations in your selected regions.
