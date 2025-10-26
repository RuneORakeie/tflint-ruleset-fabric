package rules

import (
	"fmt"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/RuneORakeie/tflint-ruleset-fabric/project"
)

// FabricWorkspaceGitCredentialsSource validates git_credentials.source values
// Valid values depend on git_provider_type:
// - GitHub: only "ConfiguredConnection"
// - AzureDevOps: "ConfiguredConnection" or "Automatic"
type FabricWorkspaceGitCredentialsSource struct {
	tflint.DefaultRule
}

func NewFabricWorkspaceGitCredentialsSource() *FabricWorkspaceGitCredentialsSource {
	return &FabricWorkspaceGitCredentialsSource{}
}

func (r *FabricWorkspaceGitCredentialsSource) Name() string {
	return "fabric_workspace_git_credentials_source"
}

func (r *FabricWorkspaceGitCredentialsSource) Enabled() bool {
	return true
}

func (r *FabricWorkspaceGitCredentialsSource) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *FabricWorkspaceGitCredentialsSource) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *FabricWorkspaceGitCredentialsSource) Check(runner tflint.Runner) error {
	resourceContent, err := runner.GetResourceContent("fabric_workspace_git", &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "git_provider_details",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "git_provider_type"},
					},
				},
			},
			{
				Type: "git_credentials",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "source"},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resourceContent.Blocks {
		// Get git_provider_type
		var providerType string
		providerBlocks := resource.Body.Blocks.OfType("git_provider_details")
		if len(providerBlocks) > 0 {
			if attr, exists := providerBlocks[0].Body.Attributes["git_provider_type"]; exists && attr.Expr != nil {
				_ = runner.EvaluateExpr(attr.Expr, &providerType, nil)
			}
		}

		// Get git_credentials.source
		credentialBlocks := resource.Body.Blocks.OfType("git_credentials")
		if len(credentialBlocks) > 0 {
			if attr, exists := credentialBlocks[0].Body.Attributes["source"]; exists && attr.Expr != nil {
				var source string
				if err := runner.EvaluateExpr(attr.Expr, &source, nil); err == nil && source != "" {
					// Validate based on provider type
					var validSources []string
					var isValid bool

					switch providerType {
					case "GitHub":
						validSources = []string{"ConfiguredConnection"}
						isValid = source == "ConfiguredConnection"
					case "AzureDevOps":
						validSources = []string{"ConfiguredConnection", "Automatic"}
						isValid = source == "ConfiguredConnection" || source == "Automatic"
					default:
						// Unknown provider type, skip validation
						continue
					}

					if !isValid {
						runner.EmitIssue(
							r,
							fmt.Sprintf("Invalid git_credentials.source '%s' for git_provider_type '%s'. Must be one of: %s",
								source, providerType, strings.Join(validSources, ", ")),
							attr.Range,
						)
					}
				}
			}
		}
	}

	return nil
}