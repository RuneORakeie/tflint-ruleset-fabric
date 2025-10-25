# fabric_deployment_pipeline_stages_display_name_length

Validates that deployment pipeline stage display names do not exceed 256 characters.

## Example

```hcl
resource "fabric_deployment_pipeline" "valid" {
  display_name = "CI/CD Pipeline"
  
  stages {
    display_name = "Development Environment" # Valid - under 256 chars
    workspace_id = fabric_workspace.dev.id
  }
  
  stages {
    display_name = "Production Environment" # Valid
    workspace_id = fabric_workspace.prod.id
  }
}

resource "fabric_deployment_pipeline" "invalid" {
  display_name = "Pipeline"
  
  stages {
    display_name = "This is an extremely long stage name that exceeds the maximum allowed length of 256 characters which will cause validation to fail and this is just an example to demonstrate what happens when you try to use a display name that is way too long for a deployment pipeline stage in Microsoft Fabric and it keeps going and going..." # Error - exceeds 256 chars
    workspace_id = fabric_workspace.dev.id
  }
}
```

## Why

Microsoft Fabric enforces a maximum length of 256 characters for deployment pipeline stage display names. This limitation ensures:

- **UI Compatibility**: Names fit properly in the Fabric portal interface
- **Database Constraints**: Alignment with backend storage limitations
- **Readability**: Encourages concise, meaningful stage names
- **API Compatibility**: Prevents issues with REST API calls

Exceeding this limit will cause the deployment pipeline creation or update to fail.

## Validation Rules

- Maximum length: 256 characters

## How to Fix

Shorten the stage display name to 256 characters or fewer:

```hcl
resource "fabric_deployment_pipeline" "example" {
  display_name = "Release Pipeline"
  
  stages {
    display_name = "Dev"  # Concise and clear
    workspace_id = fabric_workspace.dev.id
  }
  
  stages {
    display_name = "Test"
    workspace_id = fabric_workspace.test.id
  }
  
  stages {
    display_name = "Production"
    workspace_id = fabric_workspace.prod.id
  }
}
```

Use descriptions for additional context instead of very long display names:

```hcl
resource "fabric_deployment_pipeline" "example" {
  display_name = "Sales Analytics Pipeline"
  description  = "Deployment pipeline for sales analytics workspace containing customer data, revenue reports, and forecasting models"
  
  stages {
    display_name = "Development"
    description  = "Development environment for testing new features"
    workspace_id = fabric_workspace.dev.id
  }
}
```

## Configuration

```hcl
rule "fabric_deployment_pipeline_stages_display_name_length" {
  enabled = true
}
```

## Attributes

| Name | Enabled | Severity | 
|------|---------|----------|
| fabric_deployment_pipeline_stages_display_name_length | true | error |
