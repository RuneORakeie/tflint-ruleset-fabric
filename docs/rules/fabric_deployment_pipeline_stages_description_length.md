# fabric_deployment_pipeline_stages_description_length

Validates that deployment pipeline stage descriptions do not exceed 1024 characters.

## Example

```hcl
resource "fabric_deployment_pipeline" "valid" {
  display_name = "CI/CD Pipeline"
  
  stages {
    display_name = "Development"
    description  = "Development environment for testing new features and experiments" # Valid - under 1024 chars
    workspace_id = fabric_workspace.dev.id
  }
  
  stages {
    display_name = "Production"
    description  = "Production environment serving live business users"
    workspace_id = fabric_workspace.prod.id
  }
}

resource "fabric_deployment_pipeline" "invalid" {
  display_name = "Pipeline"
  
  stages {
    display_name = "Development"
    description  = "This is an extremely long description that contains way too much text and exceeds the 1024 character limit imposed by Microsoft Fabric for deployment pipeline stage descriptions and this continues with more unnecessary text to demonstrate what happens when descriptions are too verbose and contain redundant information that should probably be stored elsewhere like in documentation or wiki pages instead of being embedded directly in the stage description field..." # Error - exceeds 1024 chars
    workspace_id = fabric_workspace.dev.id
  }
}
```

## Why

Microsoft Fabric enforces a maximum length of 1024 characters for deployment pipeline stage descriptions. This limitation ensures:

- **Database Constraints**: Alignment with backend storage requirements
- **Performance**: Prevents excessive data transfer in API responses
- **UI Compatibility**: Ensures descriptions display properly in the portal
- **Best Practices**: Encourages concise, focused documentation

Exceeding this limit will cause the deployment pipeline creation or update to fail.

## Validation Rules

- Maximum length: 1024 characters

## How to Fix

Shorten the stage description to 1024 characters or fewer. Focus on essential information:

```hcl
resource "fabric_deployment_pipeline" "example" {
  display_name = "Analytics Pipeline"
  
  stages {
    display_name = "Development"
    description  = "Development environment for data analysts and engineers. Used for testing new reports, datasets, and ETL processes before promoting to production."
    workspace_id = fabric_workspace.dev.id
  }
  
  stages {
    display_name = "UAT"
    description  = "User Acceptance Testing environment. Business users validate changes before production deployment."
    workspace_id = fabric_workspace.uat.id
  }
  
  stages {
    display_name = "Production"
    description  = "Production environment serving 500+ business users. Contains certified reports and datasets."
    workspace_id = fabric_workspace.prod.id
  }
}
```

For extensive documentation, consider maintaining it externally and linking to it:

```hcl
stages {
  display_name = "Production"
  description  = "Production environment for sales analytics. See full documentation: https://wiki.company.com/fabric/sales-analytics-prod"
  workspace_id = fabric_workspace.prod.id
}
```

## Configuration

```hcl
rule "fabric_deployment_pipeline_stages_description_length" {
  enabled = true
}
```

## Attributes

| Name | Enabled | Severity | 
|------|---------|----------|
| fabric_deployment_pipeline_stages_description_length | true | error |
