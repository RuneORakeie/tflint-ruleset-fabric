# Schema Generation

This directory is used to generate the Terraform provider schema.

## Generate Schema

1. Initialize Terraform:
```bash
terraform init
```

2. Generate schema JSON:
```bash
terraform providers schema -json > schema.json
```

## Notes

- The `schema.json` file is used by the rule generator to validate that attributes exist in the Terraform provider
- Regenerate this file when the provider is updated
