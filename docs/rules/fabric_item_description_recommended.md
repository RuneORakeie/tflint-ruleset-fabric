# fabric_item_description_recommended

Recommends adding descriptions to Fabric items to improve documentation and governance.

## Example

```hcl
resource "fabric_lakehouse" "valid" {
  workspace_id = fabric_workspace.example.id
  display_name = "Sales Data"
  description  = "Lakehouse containing sales transactions and customer data for analytics" # Good
}

resource "fabric_notebook" "no_description" {
  workspace_id = fabric_workspace.example.id
  display_name = "Analysis Notebook"
  # Missing description - will emit warning
}

resource "fabric_warehouse" "empty_description" {
  workspace_id = fabric_workspace.example.id
  display_name = "Data Warehouse"
  description  = "" # Empty description - will emit warning
}
```

## Why

Descriptions help document your Fabric environment, making it easier for teams to understand the purpose and ownership of resources, especially in large collaborative environments.

**Benefits of adding descriptions:**
- **Documentation**: Clear purpose and context for each resource
- **Governance**: Better understanding of resource ownership and responsibilities
- **Onboarding**: Helps new team members navigate the environment
- **Compliance**: Meets organizational documentation requirements
- **Searchability**: Makes it easier to find relevant resources

## Applies To

This rule checks descriptions for the following resource types:
- `fabric_connection`
- `fabric_deployment_pipeline`
- `fabric_domain`
- `fabric_eventhouse`
- `fabric_kql_database`
- `fabric_kql_queryset`
- `fabric_lakehouse`
- `fabric_ml_experiment`
- `fabric_ml_model`
- `fabric_notebook`
- `fabric_report`
- `fabric_semantic_model`
- `fabric_spark_job_definition`
- `fabric_warehouse`
- `fabric_workspace`

## How to Fix

Add meaningful descriptions that include:
- The purpose of the resource
- The owner or team responsible
- Any relevant business context
- Data sources or dependencies (if applicable)

```hcl
resource "fabric_lakehouse" "sales" {
  workspace_id = fabric_workspace.analytics.id
  display_name = "Sales Lakehouse"
  description  = "Central lakehouse for sales data. Contains customer transactions, product catalog, and sales metrics. Owner: sales-analytics-team@company.com. Updated nightly from CRM system."
}
```

## Configuration

```hcl
rule "fabric_item_description_recommended" {
  enabled = true
}
```

## Attributes

| Name | Enabled | Severity | 
|------|---------|----------|
| fabric_item_description_recommended | true | warning |
