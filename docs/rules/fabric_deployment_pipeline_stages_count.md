# fabric_deployment_pipeline_stages_count

Validates that deployment pipelines have between 2 and 10 stages as per Microsoft Fabric requirements.

## Example

```hcl
# Valid - has 3 stages (within 2-10 range)
resource "fabric_deployment_pipeline" "valid" {
  display_name = "CI/CD Pipeline"
  
  stages {
    display_name = "Development"
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

# Invalid - only 1 stage (minimum is 2)
resource "fabric_deployment_pipeline" "too_few" {
  display_name = "Invalid Pipeline"
  
  stages {
    display_name = "Production"
    workspace_id = fabric_workspace.prod.id
  }
  # Error: Must have at least 2 stages
}
```

## Why

Microsoft Fabric deployment pipelines require a minimum of 2 stages and support a maximum of 10 stages. This limitation is enforced by the Fabric service:

- **Minimum (2 stages)**: Deployment pipelines are designed for promoting content between environments (e.g., Dev → Prod), which requires at least two stages
- **Maximum (10 stages)**: This is a platform limitation to maintain manageable deployment workflows

## Validation Rules

- Minimum stages: 2
- Maximum stages: 10

## How to Fix

Ensure your deployment pipeline has between 2 and 10 stages:

```hcl
resource "fabric_deployment_pipeline" "example" {
  display_name = "Standard Pipeline"
  
  stages {
    display_name = "Development"
    workspace_id = fabric_workspace.dev.id
  }
  
  stages {
    display_name = "Production"
    workspace_id = fabric_workspace.prod.id
  }
}
```

For more complex scenarios, you can add additional stages:

```hcl
resource "fabric_deployment_pipeline" "complex" {
  display_name = "Multi-Stage Pipeline"
  
  stages {
    display_name = "Development"
    workspace_id = fabric_workspace.dev.id
  }
  
  stages {
    display_name = "Integration"
    workspace_id = fabric_workspace.int.id
  }
  
  stages {
    display_name = "UAT"
    workspace_id = fabric_workspace.uat.id
  }
  
  stages {
    display_name = "Production"
    workspace_id = fabric_workspace.prod.id
  }
}
```

## Configuration

```hcl
rule "fabric_deployment_pipeline_stages_count" {
  enabled = true
}
```

## Attributes

| Name | Enabled | Severity | 
|------|---------|----------|
| fabric_deployment_pipeline_stages_count | true | error |
