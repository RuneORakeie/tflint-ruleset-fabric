# Fabric TFLint Ruleset - Mapping Files

This directory contains mapping files that connect Terraform Fabric provider resources to their REST API specifications. These mappings enable automatic generation of validation rules.

## Important: Workspace ID vs Folder ID

**Critical Understanding**: In the Microsoft Fabric REST API, the parameter often named `folderId` is actually the workspace ID. This is a key difference from other Microsoft services.

```
API Path Example: /v1/workspaces/{folderId}/items/{itemId}
Terraform Attribute: workspace_id

The folderId parameter in the API = workspace_id attribute in Terraform
```

## File Structure

Each mapping file follows this pattern:

```hcl
mapping "fabric_<resource_name>" {
  import_path = "fabric-rest-api-specs/specification/<service>/v1/<file>.json"
  
  // Map Terraform attributes to API spec definitions
  terraform_attribute = ApiDefinitionName
  nested_attribute = ParentObject.childProperty
  complex_attribute = any  // Skip validation for complex types
}
```

## Current Mapping Files

### Core Resources
- `fabric_workspace.hcl` - Workspace configuration
- `fabric_lakehouse.hcl` - Lakehouse data storage
- `fabric_warehouse.hcl` - Data warehouse
- `fabric_kql_database.hcl` - KQL database in Eventhouse

### Compute Resources
- `fabric_spark_job_definition.hcl` - Spark job definitions
- `fabric_notebook.hcl` - Notebook resources

### Data Pipeline Resources
- `fabric_data_pipeline.hcl` - Data pipeline definitions
- `fabric_deployment_pipeline.hcl` - Deployment pipeline with stages

### Analytics Resources
- `fabric_eventhouse.hcl` - Real-time analytics eventhouse
- `fabric_semantic_model.hcl` - Power BI semantic models

## How to Use

### 1. Prerequisites

Before using these mappings, you need:

1. **Fabric REST API Specifications** in OpenAPI/Swagger JSON format
   - These should be organized in a `fabric-rest-api-specs` directory
   - Structure: `specification/<service>/v1/<resource>.json`

2. **Terraform Provider Schema** 
   - The `schema.json` file in the project root contains the provider schema
   - Generated from the terraform-provider-fabric

### 2. Adjusting Import Paths

The example mappings use placeholder paths. You'll need to update the `import_path` in each file to match your actual API spec structure:

```hcl
// Example - update this path
import_path = "fabric-rest-api-specs/specification/lakehouse/v1/lakehouses.json"

// To match your actual path, for example:
import_path = "azure-rest-api-specs/specification/fabric/lakehouse/stable/2024-05-01/lakehouse.json"
```

### 3. Running the Generator

From the tools/apispec-rule-gen directory:

```bash
go run main.go \
  --base-path=/path/to/apispec-rule-gen \
  --rules-path=/path/to/rules \
  --docs-path=/path/to/docs
```

The generator will:
1. Parse all `.hcl` files in the mappings directory
2. Load the corresponding API spec JSON files
3. Generate validation rules for attributes with:
   - Enum values (string validation)
   - Patterns (regex validation)
   - Min/max values (integer validation)
4. Create Go rule files and documentation

## Mapping Patterns

### Simple Attribute
```hcl
display_name = ItemDisplayName
```
Maps the `display_name` Terraform attribute to the `ItemDisplayName` definition in the API spec.

### Nested Property
```hcl
description = ItemProperties.description
```
Maps to a property within a parent object in the API spec.

### Parameter Reference
```hcl
workspace_id = folderIdParameter
```
Maps to a parameter defined in the API spec's `parameters` section. Note: `folderId` in Fabric API is the workspace ID.

### Complex Type (Skip Validation)
```hcl
definition = any
```
Use `any` for complex objects that don't need validation (like JSON payloads).

### Block with Children
```hcl
stages = {
  display_name = StageProperties.displayName
  workspace_id = StageProperties.workspaceId
}
```
Maps a block (nested resource) with multiple child attributes.

## Common Fabric Patterns

### Workspace-Scoped Resources

Most Fabric items exist within a workspace:

```hcl
mapping "fabric_<item>" {
  import_path = "..."
  
  display_name = ItemName
  workspace_id = folderIdParameter  # CRITICAL: folderId = workspace_id
  description = ItemProperties.description
}
```

### Workspace Configuration

Workspace-level settings:

```hcl
mapping "fabric_workspace" {
  import_path = "..."
  
  display_name = WorkspaceName
  capacity_id = WorkspaceProperties.capacityId
  description = WorkspaceProperties.description
}
```

### Items with Definitions

Many Fabric items have complex `definition` payloads:

```hcl
mapping "fabric_notebook" {
  import_path = "..."
  
  display_name = NotebookName
  workspace_id = folderIdParameter
  definition = any  # Complex notebook JSON structure
}
```

## Troubleshooting

### Error: "resource not found in the Terraform schema"
- Verify the resource name matches exactly (check `schema.json`)
- Ensure you're using the correct resource type (e.g., `fabric_lakehouse` not `fabric_lake_house`)

### Error: "attribute not found in the Terraform schema"
- Check attribute name spelling and casing
- Verify the attribute exists in the provider schema
- Some attributes might be in nested blocks

### Error: "is expected as string, but not"
- Type mismatch between API spec and Terraform schema
- Check if the API spec defines the field as the correct type
- Consider if the field should be `string` vs `number`

### No Rules Generated
- Verify the API spec has validation constraints:
  - `enum` for string values
  - `pattern` for regex validation
  - `minimum`/`maximum` for integer values
- Check that the `import_path` points to a valid, accessible JSON file
- Ensure the API spec follows OpenAPI/Swagger format

### Error: "Unable to open API spec file"
- Verify the `import_path` is correct relative to the generator's working directory
- Check file permissions
- Ensure the fabric-rest-api-specs directory structure matches your paths

## API Spec Requirements

For the generator to create validation rules, the API spec must contain:

### Enum Validation (Strings)
```json
{
  "type": "string",
  "enum": ["Development", "Test", "Production"]
}
```

### Pattern Validation (Strings)
```json
{
  "type": "string",
  "pattern": "^[a-zA-Z0-9_-]+$",
  "minLength": 1,
  "maxLength": 256
}
```

### Range Validation (Integers)
```json
{
  "type": "integer",
  "minimum": 1,
  "maximum": 100
}
```

## Next Steps

1. **Obtain or create Fabric REST API specs** in OpenAPI format
   - Work with the fabric-rest-api-specs repository
   - Or generate from Fabric API documentation

2. **Update import paths** in the mapping files
   - Match your actual API spec file locations

3. **Add missing mappings** for additional Fabric resources
   - Use the examples as templates
   - Focus on attributes that have validation needs

4. **Test the generator**
   - Run the tool and verify rules are created correctly
   - Check generated Go files in the rules directory
   - Review generated documentation

5. **Iterate and refine**
   - Add more detailed mappings as you understand the API structure
   - Document any Fabric-specific quirks or patterns

## Contributing

When adding new mapping files:

1. Follow the naming convention: `fabric_<resource_type>.hcl`
2. Add comments explaining Fabric-specific behaviors
3. Document any workspace_id / folder_id mappings
4. Test that rules generate successfully
5. Update this README with the new mapping

## Resources

- [TFLint Ruleset Azure RM](https://github.com/terraform-linters/tflint-ruleset-azurerm) - Reference implementation
- [Terraform Provider Fabric](https://github.com/microsoft/terraform-provider-fabric) - Provider source
- [Microsoft Fabric REST API](https://learn.microsoft.com/en-us/rest/api/fabric/) - API documentation
