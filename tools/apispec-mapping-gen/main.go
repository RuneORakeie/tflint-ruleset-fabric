package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

// Tool to analyze Fabric REST API specs and generate mapping files
// This scans the fabric-rest-api-specs directory and extracts validation constraints
// directoryPrefixes maps API spec directory names to resource name prefixes
// Use this when resources in a directory should have a specific prefix
var directoryPrefixes = map[string]string{
	"spark": "spark_",  // spark/customPool -> fabric_spark_custom_pool
}
// resourceNameOverrides maps API spec resource names to Terraform resource names
// Use this when the API spec name differs from the actual Terraform resource name
var resourceNameOverrides = map[string]string{
	"graphqlapi":                     "graphql_api",  // GraphQL naming
	"kqldatabase":                    "kql_database",  // KQL naming
	"managed_private_endpoint":       "workspace_managed_private_endpoint", // MPE in platform folder
	"mirroredazuredatabrickscatalog": "mirrored_azure_databricks_catalog",  // Long name fix
	"mlmodel":                        "ml_model",  // ML naming
	"mlexperiment":                   "ml_experiment",  // ML naming
	"reflex":                         "activator",  // Reflex was renamed to Activator
	"sparkjobdefinition":             "spark_job_definition",  // Naming convention fix
	"sqldatabase":                    "sql_database",  // SQL naming
}

// resourceMergeRules defines resources that need properties from multiple Create*Request types merged
// Key: base resource name (e.g., "connection")
// Value: list of additional Create*Request types to merge properties from
var resourceMergeRules = map[string][]string{
	"connection": {
		"CreateCloudConnectionRequest",                 // For ShareableCloud connectivity_type
		"CreateVirtualNetworkGatewayConnectionRequest", // For VirtualNetworkGateway connectivity_type
		"CreateOnPremisesConnectionRequest",            // For OnPremises connectivity_type (not yet in Terraform)
	},
	"gateway": {
		"CreateVirtualNetworkGatewayRequest", // For VirtualNetwork gateway type
		"CreateOnPremisesGatewayRequest",     // For OnPremises gateway type
	},
}

// excludedRequestTypes lists Create*Request types that should NOT generate their own mappings
// because they are merged into another resource via resourceMergeRules
var excludedRequestTypes = map[string]bool{
	"CreateCloudConnectionRequest":                 true, // Merged into fabric_connection
	"CreateVirtualNetworkGatewayConnectionRequest": true, // Merged into fabric_connection
	"CreateOnPremisesConnectionRequest":            true, // Not supported in Terraform at all
	"CreateCredentialDetailsRequest":               true, // Nested object, not a resource
	"CreateVirtualNetworkGatewayRequest":           true, // Merged into fabric_gateway (VirtualNetwork type)
	"CreateOnPremisesGatewayRequest":               true, // Merged into fabric_gateway (OnPremises type)
}

type SwaggerSpec struct {
	Swagger     string                        `json:"swagger"`
	Info        map[string]interface{}        `json:"info"`
	Paths       map[string]interface{}        `json:"paths"`
	Definitions map[string]DefinitionSchema   `json:"definitions"`
}

type DefinitionSchema struct {
	Description string                     `json:"description"`
	Type        string                     `json:"type"`
	Required    []string                   `json:"required"`
	Properties  map[string]PropertySchema  `json:"properties"`
	AllOf       []map[string]interface{}   `json:"allOf"`
	ReadOnly    bool                       `json:"readOnly"`
	// Fields for simple type definitions (like enums)
	Format      string                     `json:"format"`
	MaxLength   int                        `json:"maxLength"`
	MinLength   int                        `json:"minLength"`
	Pattern     string                     `json:"pattern"`
	Enum        []string                   `json:"enum"`
}

type PropertySchema struct {
	Description string      `json:"description"`
	Type        string      `json:"type"`
	Format      string      `json:"format"`
	MaxLength   int         `json:"maxLength"`
	MinLength   int         `json:"minLength"`
	Pattern     string      `json:"pattern"`
	Enum        []string    `json:"enum"`
	ReadOnly    bool        `json:"readOnly"`
	Ref         string      `json:"$ref"`
}

type ResourceInfo struct {
	SpecPath          string
	CreateRequestType string
	UpdateRequestType string
	Constraints       map[string]PropertyConstraints
}

type PropertyConstraints struct {
	MaxLength      int
	MinLength      int
	Pattern        string
	Format         string
	Enum           []string
	Description    string
	Required       bool
	SourceRequests []string // Tracks all Create*Request types this property came from
	ApiRef         string   // For manual attributes, stores the original api_ref
}

// Structures for parsing existing HCL mapping files
type existingMappingFile struct {
	Mappings []existingMapping `hcl:"mapping,block"`
}

type existingMapping struct {
	Resource   string                     `hcl:"resource,label"`
	ImportPath string                     `hcl:"import_path"`
	Attributes []existingAttributeMapping `hcl:"attribute,block"`
}

type existingAttributeMapping struct {
	Name         string   `hcl:"name,label"`
	ApiRef       string   `hcl:"api_ref"`
	MaxLength    *int     `hcl:"max_length,optional"`
	MinLength    *int     `hcl:"min_length,optional"`
	Pattern      *string  `hcl:"pattern,optional"`
	WarnOnExceed *bool    `hcl:"warn_on_exceed,optional"`
	ValidValues  []string `hcl:"valid_values,optional"`
}

func main() {
	var specsPath, outputPath string
	var skipExisting bool
	flag.StringVar(&specsPath, "specs", "../fabric-rest-api-specs", "Path to fabric-rest-api-specs directory")
	flag.StringVar(&outputPath, "output", "mappings", "Output directory for mapping files")
	flag.BoolVar(&skipExisting, "skip-existing", false, "Skip files that already exist (default: false, will merge updates)")
	flag.Parse()

	fmt.Println("Analyzing Fabric API specs from:", specsPath)

	// Scan all spec directories
	entries, err := os.ReadDir(specsPath)
	if err != nil {
		fmt.Printf("Error reading specs directory: %v\n", err)
		os.Exit(1)
	}

	resourceMap := make(map[string]*ResourceInfo)

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		dirName := entry.Name()
		
		// Skip common directory
		if dirName == "common" || dirName == ".git" {
			continue
		}

		// Look for definition files in the main directory
		var definitionFiles []string
		
		// Check for common naming patterns in main directory
		possibleFiles := []string{"definitions.json", dirName + ".json", "swagger.json"}
		for _, filename := range possibleFiles {
			testPath := filepath.Join(specsPath, dirName, filename)
			if _, err := os.Stat(testPath); err == nil {
				// Found a definition file - check if it has Swagger definitions
				data, err := os.ReadFile(testPath)
				if err != nil {
					continue
				}
				
				var spec SwaggerSpec
				if err := json.Unmarshal(data, &spec); err != nil {
					continue
				}
				
				// Only use if it has definitions
				if len(spec.Definitions) > 0 {
					definitionFiles = append(definitionFiles, testPath)
				}
			}
		}
		
		// Also check for a definitions/ subdirectory
		defsSubdir := filepath.Join(specsPath, dirName, "definitions")
		if stat, err := os.Stat(defsSubdir); err == nil && stat.IsDir() {
			subFiles, err := os.ReadDir(defsSubdir)
			if err == nil {
				for _, subFile := range subFiles {
					if !subFile.IsDir() && strings.HasSuffix(subFile.Name(), ".json") {
						testPath := filepath.Join(defsSubdir, subFile.Name())
						data, err := os.ReadFile(testPath)
						if err != nil {
							continue
						}
						
						var spec SwaggerSpec
						if err := json.Unmarshal(data, &spec); err != nil {
							continue
						}
						
						// Only use if it has definitions
						if len(spec.Definitions) > 0 {
							definitionFiles = append(definitionFiles, testPath)
						}
					}
				}
			}
		}
		
		if len(definitionFiles) == 0 {
			continue
		}

		fmt.Printf("\nAnalyzing %s (found %d definition files)...\n", dirName, len(definitionFiles))
		
		// Analyze all definition files for this directory
		for _, defPath := range definitionFiles {
			data, err := os.ReadFile(defPath)
			if err != nil {
				fmt.Printf("  Error reading %s: %v\n", defPath, err)
				continue
			}

			var spec SwaggerSpec
			if err := json.Unmarshal(data, &spec); err != nil {
				fmt.Printf("  Error parsing %s: %v\n", defPath, err)
				continue
			}

			// Find Create/Update request definitions - there may be multiple in one file
			relPath, _ := filepath.Rel(specsPath, defPath)
			resources := analyzeSpecForResources(dirName, relPath, spec)
			for resourceName, info := range resources {
				resourceMap[resourceName] = info
				fmt.Printf("  ✓ Found %s with %d constraints in %s\n", resourceName, len(info.Constraints), filepath.Base(defPath))
			}
		}
	}

	fmt.Printf("\n✅ Analyzed %d resource types\n", len(resourceMap))
	
	// Generate mapping files
	if err := generateMappingFiles(resourceMap, outputPath, skipExisting); err != nil {
		fmt.Printf("Error generating mapping files: %v\n", err)
		os.Exit(1)
	}
	
	// Generate summary report
	generateSummaryReport(resourceMap)
}

func analyzeSpecForResources(dirName string, relPath string, spec SwaggerSpec) map[string]*ResourceInfo {
	resources := make(map[string]*ResourceInfo)

	// Look for Create*Request definitions - there may be multiple in one file
	for defName, defSchema := range spec.Definitions {
		if strings.HasPrefix(defName, "Create") && strings.HasSuffix(defName, "Request") {
			// Skip request types that are merged into other resources
			if excludedRequestTypes[defName] {
				fmt.Printf("    ⊘ Skipping %s (merged into another resource)\n", defName)
				continue
			}
			
			// Extract resource name from CreateXxxRequest
			// e.g., CreateTagsRequest -> tags, CreateDomainRequest -> domain
			resourceName := strings.TrimSuffix(strings.TrimPrefix(defName, "Create"), "Request")
			resourceName = toSnakeCase(resourceName)
			
			// Check if there's an override for this resource name
			if override, exists := resourceNameOverrides[resourceName]; exists {
				fmt.Printf("    → Mapping %s to %s (name override)\n", resourceName, override)
				resourceName = override
			}
			
			// Apply directory-specific prefix if configured
			if prefix, exists := directoryPrefixes[dirName]; exists {
				fmt.Printf("    → Adding prefix '%s' to %s (directory: %s)\n", prefix, resourceName, dirName)
				resourceName = prefix + resourceName
			}
			
			info := &ResourceInfo{
				SpecPath:          relPath,
				CreateRequestType: defName,
				Constraints:       make(map[string]PropertyConstraints),
			}
			
			extractConstraints(defSchema, info.Constraints, defName, &spec)
			
			// Also look for corresponding Update request
			updateDefName := "Update" + strings.TrimPrefix(defName, "Create")
			if updateSchema, exists := spec.Definitions[updateDefName]; exists {
				info.UpdateRequestType = updateDefName
				// Extract Update constraints but don't add them if they duplicate Create properties
				// We'll merge the constraints but keep the Create api_ref
				extractConstraints(updateSchema, info.Constraints, updateDefName, &spec)
			}
			
			// Check if this resource needs properties merged from other request types
			if mergeTypes, shouldMerge := resourceMergeRules[resourceName]; shouldMerge {
				fmt.Printf("    → Merging properties from %d additional request types for %s\n", len(mergeTypes), resourceName)
				for _, mergeType := range mergeTypes {
					if mergeSchema, exists := spec.Definitions[mergeType]; exists {
						fmt.Printf("      ✓ Merging %s\n", mergeType)
						extractConstraints(mergeSchema, info.Constraints, mergeType, &spec)
					} else {
						fmt.Printf("      ⚠️  Merge type %s not found in spec\n", mergeType)
					}
				}
			}
			
			resources[resourceName] = info
		}
	}

	return resources
}

func extractConstraints(schema DefinitionSchema, constraints map[string]PropertyConstraints, sourceRequest string, spec *SwaggerSpec) {
	for propName, propSchema := range schema.Properties {
		// Skip read-only properties
		if propSchema.ReadOnly {
			continue
		}

		constraint := PropertyConstraints{
			MaxLength:      propSchema.MaxLength,
			MinLength:      propSchema.MinLength,
			Pattern:        propSchema.Pattern,
			Format:         propSchema.Format,
			Enum:           propSchema.Enum,
			Description:    propSchema.Description,
			SourceRequests: []string{sourceRequest},
		}
		
		// If this property has a $ref, resolve it to get enum values
		if propSchema.Ref != "" && len(constraint.Enum) == 0 {
			// Extract definition name from $ref (e.g., "#/definitions/ConnectivityType" -> "ConnectivityType")
			refParts := strings.Split(propSchema.Ref, "/")
			if len(refParts) > 0 {
				refDefName := refParts[len(refParts)-1]
				if refDef, exists := spec.Definitions[refDefName]; exists {
					// Use enum values from the referenced definition
					if len(refDef.Enum) > 0 {
						constraint.Enum = refDef.Enum
					}
					// Also get other constraints from the ref if not already set
					if constraint.MaxLength == 0 && refDef.MaxLength > 0 {
						constraint.MaxLength = refDef.MaxLength
					}
					if constraint.MinLength == 0 && refDef.MinLength > 0 {
						constraint.MinLength = refDef.MinLength
					}
					if constraint.Pattern == "" && refDef.Pattern != "" {
						constraint.Pattern = refDef.Pattern
					}
					if constraint.Format == "" && refDef.Format != "" {
						constraint.Format = refDef.Format
					}
				}
			}
		}
		
		// Check if required
		for _, req := range schema.Required {
			if req == propName {
				constraint.Required = true
				break
			}
		}

		// Extract maxLength from description using multiple patterns
		if constraint.MaxLength == 0 && propSchema.Description != "" {
			constraint.MaxLength = extractMaxLengthFromDescription(propSchema.Description)
		}

		// Only add if there are actual constraints
		if constraint.MaxLength > 0 || constraint.MinLength > 0 || constraint.Pattern != "" || 
		   constraint.Format != "" || len(constraint.Enum) > 0 || constraint.Required {
			
			// For properties that might appear in multiple merged types,
			// create a unique key that includes the source request type
			// This prevents overwriting when merging multiple request types
			uniqueKey := propName + ":" + sourceRequest
			
			// If this is from an Update request, check if a Create request version exists
			// and skip if so (we prefer Create for api_ref)
			if strings.HasPrefix(sourceRequest, "Update") {
				createRequest := "Create" + strings.TrimPrefix(sourceRequest, "Update")
				createKey := propName + ":" + createRequest
				if _, exists := constraints[createKey]; exists {
					// Skip this Update property - we already have it from Create
					continue
				}
			}
			
			constraints[uniqueKey] = constraint
		}
	}
}

// extractMaxLengthFromDescription extracts max length from description text using various patterns
func extractMaxLengthFromDescription(description string) int {
	patterns := []string{
		`String length must be at most (\d+)`,
		`length must not exceed (\d+)`,
		`Maximum length is (\d+)`,
		`cannot contain more than (\d+) characters`,
		`cannot exceed (\d+) characters`,
		`must be (\d+) characters or less`,
		`up to (\d+) characters`,
		`max(?:imum)?\s+(\d+)\s+char(?:acter)?s?`,
	}

	for _, pattern := range patterns {
		re := regexp.MustCompile(`(?i)` + pattern) // case-insensitive
		if matches := re.FindStringSubmatch(description); len(matches) > 1 {
			if length, err := strconv.Atoi(matches[1]); err == nil {
				return length
			}
		}
	}

	return 0
}

func generateMappingFiles(resourceMap map[string]*ResourceInfo, outputPath string, skipExisting bool) error {
	if err := os.MkdirAll(outputPath, 0755); err != nil {
		return err
	}

	// Sort resource names for consistent output
	var resourceNames []string
	for name := range resourceMap {
		resourceNames = append(resourceNames, name)
	}
	sort.Strings(resourceNames)

	created, skipped, updated := 0, 0, 0
	
	for _, resourceName := range resourceNames {
		info := resourceMap[resourceName]
		
		// Resource name is already in snake_case from analyzeSpecForResources
		// Just add the fabric_ prefix
		tfResourceName := "fabric_" + resourceName
		
		status, err := generateMappingFile(tfResourceName, resourceName, info, outputPath, skipExisting)
		if err != nil {
			return fmt.Errorf("error generating mapping for %s: %v", resourceName, err)
		}
		
		switch status {
		case "created":
			created++
		case "skipped":
			skipped++
		case "updated":
			updated++
		}
	}

	fmt.Printf("\n✅ Mapping files: %d created, %d skipped, %d updated in %s\n", 
		created, skipped, updated, outputPath)
	return nil
}

// extractRequestTypeSuffix extracts a meaningful suffix from a request type name
// to create unique attribute names for merged properties
func extractRequestTypeSuffix(requestType string) string {
	// Map of known request types to their suffixes
	suffixMap := map[string]string{
		"CreateCloudConnectionRequest":                 "cloud",
		"CreateVirtualNetworkGatewayConnectionRequest": "vnet",
		"CreateOnPremisesConnectionRequest":            "onprem",
		"UpdateConnectionRequest":                      "update",
	}
	
	if suffix, exists := suffixMap[requestType]; exists {
		return suffix
	}
	
	// Fallback: extract a generic suffix from the request type name
	// e.g., "CreateFooBarRequest" -> "foo_bar"
	name := strings.TrimPrefix(requestType, "Create")
	name = strings.TrimPrefix(name, "Update")
	name = strings.TrimSuffix(name, "Request")
	return toSnakeCase(name)
}

// parseExistingMapping parses an existing HCL mapping file and returns the attributes
func parseExistingMapping(filename string) (map[string]*existingAttributeMapping, error) {
	parser := hclparse.NewParser()
	f, diags := parser.ParseHCLFile(filename)
	if diags.HasErrors() {
		return nil, fmt.Errorf("failed to parse HCL: %v", diags)
	}

	var mf existingMappingFile
	diags = gohcl.DecodeBody(f.Body, nil, &mf)
	if diags.HasErrors() {
		return nil, fmt.Errorf("failed to decode HCL: %v", diags)
	}

	// Build map of attribute name -> attribute mapping
	existingAttrs := make(map[string]*existingAttributeMapping)
	if len(mf.Mappings) > 0 {
		for i := range mf.Mappings[0].Attributes {
			attr := &mf.Mappings[0].Attributes[i]
			existingAttrs[attr.Name] = attr
		}
	}

	return existingAttrs, nil
}

// mergeConstraints merges API spec constraints with existing manual constraints
// Manual constraints (from existing file) take precedence
func mergeConstraints(apiConstraint PropertyConstraints, existing *existingAttributeMapping) PropertyConstraints {
	merged := apiConstraint
	
	// Manual constraints override API spec
	if existing != nil {
		if existing.MaxLength != nil {
			merged.MaxLength = *existing.MaxLength
		}
		if existing.MinLength != nil {
			merged.MinLength = *existing.MinLength
		}
		if existing.Pattern != nil {
			merged.Pattern = *existing.Pattern
		}
		if len(existing.ValidValues) > 0 {
			merged.Enum = existing.ValidValues
		}
	}
	
	return merged
}

func generateMappingFile(tfResourceName, specDir string, info *ResourceInfo, outputPath string, skipExisting bool) (string, error) {
	filename := filepath.Join(outputPath, tfResourceName+".hcl")
	
	// Parse existing mapping file if it exists
	var existingAttrs map[string]*existingAttributeMapping
	fileExists := false
	if _, err := os.Stat(filename); err == nil {
		fileExists = true
		existingAttrs, err = parseExistingMapping(filename)
		if err != nil {
			fmt.Printf("  ⚠️  Warning: could not parse existing %s: %v\n", tfResourceName, err)
			if skipExisting {
				fmt.Printf("  ⚠️  Skipping %s (parse error + skipExisting=true)\n", tfResourceName)
				return "skipped", nil
			}
			// Continue with empty existing attrs if not skipping
			existingAttrs = make(map[string]*existingAttributeMapping)
		} else {
			if skipExisting {
				fmt.Printf("  ⚠️  Skipping %s (already exists and skipExisting=true)\n", tfResourceName)
				return "skipped", nil
			}
			fmt.Printf("  ♻️  Merging updates into %s (%d existing attributes)\n", tfResourceName, len(existingAttrs))
		}
	} else {
		fmt.Printf("  ✨ Creating new mapping: %s\n", tfResourceName)
		existingAttrs = make(map[string]*existingAttributeMapping)
	}
	
	// Merge constraints: existing manual constraints take precedence
	mergedConstraints := make(map[string]PropertyConstraints)
	allAttrNames := make(map[string]bool)
	
	// Add all API spec attributes
	for key, constraint := range info.Constraints {
		parts := strings.Split(key, ":")
		propName := parts[0]
		tfAttrName := toSnakeCase(propName)
		allAttrNames[tfAttrName] = true
		
		// Merge with existing if present
		merged := mergeConstraints(constraint, existingAttrs[tfAttrName])
		mergedConstraints[key] = merged
	}
	
	// Add existing attributes that aren't in the API spec (manual additions)
	for attrName, existingAttr := range existingAttrs {
		if !allAttrNames[attrName] {
			fmt.Printf("    📌 Preserving manual attribute: %s\n", attrName)
			// Create a constraint from the existing manual attribute
			manualConstraint := PropertyConstraints{
				Description: fmt.Sprintf("Manual attribute (not in API spec)"),
				ApiRef:      existingAttr.ApiRef, // Preserve the original api_ref
			}
			if existingAttr.MaxLength != nil {
				manualConstraint.MaxLength = *existingAttr.MaxLength
			}
			if existingAttr.MinLength != nil {
				manualConstraint.MinLength = *existingAttr.MinLength
			}
			if existingAttr.Pattern != nil {
				manualConstraint.Pattern = *existingAttr.Pattern
			}
			if len(existingAttr.ValidValues) > 0 {
				manualConstraint.Enum = existingAttr.ValidValues
			}
			
			// Use the attribute name itself as the key (with :manual suffix for tracking)
			mergedConstraints[attrName+":manual"] = manualConstraint
		}
	}
	
	// Update info.Constraints with merged values
	info.Constraints = mergedConstraints
	
	var content strings.Builder
	
	content.WriteString(fmt.Sprintf("// Mapping for %s resource\n", tfResourceName))
	content.WriteString(fmt.Sprintf("// Auto-generated from %s\n", info.SpecPath))
	
	// Check if this resource had properties merged from multiple request types
	resourceNameWithoutPrefix := strings.TrimPrefix(tfResourceName, "fabric_")
	if mergeTypes, wasMerged := resourceMergeRules[resourceNameWithoutPrefix]; wasMerged {
		content.WriteString(fmt.Sprintf("// Properties merged from: %s", info.CreateRequestType))
		for _, mergeType := range mergeTypes {
			content.WriteString(fmt.Sprintf(", %s", mergeType))
		}
		content.WriteString("\n")
	}
	
	content.WriteString("// DO NOT EDIT auto-generated sections directly.\n")
	content.WriteString("// Add custom constraints with // MANUAL: comment to preserve during updates.\n\n")
	
	content.WriteString(fmt.Sprintf("mapping \"%s\" {\n", tfResourceName))
	content.WriteString(fmt.Sprintf("  import_path = \"%s\"\n\n", info.SpecPath))
	
	// Sort properties for consistent output
	var propKeys []string
	for key := range info.Constraints {
		propKeys = append(propKeys, key)
	}
	sort.Strings(propKeys)
	
	// Generate attribute blocks with comments about constraints
       for _, propKey := range propKeys {
	       constraint := info.Constraints[propKey]
	       parts := strings.Split(propKey, ":")
	       propName := parts[0]
	       var sourceRequest string
	       if len(parts) > 1 {
		       sourceRequest = parts[1]
	       }
	       // For manual attributes, use the original attribute name
	       var tfAttrName string
	       if sourceRequest == "manual" {
		       tfAttrName = propKey[:len(propName)] // Use the original attribute name
	       } else {
		       tfAttrName = toSnakeCase(propName)
	       }
	       // If this property comes from a merged request type (not the base CreateRequest),
	       // append a suffix to make the attribute name unique
	       if sourceRequest != "" && sourceRequest != "manual" && sourceRequest != info.CreateRequestType {
		       suffix := extractRequestTypeSuffix(sourceRequest)
		       if suffix != "" {
			       tfAttrName += "_" + suffix
		       }
	       }
	       // Build comment describing the attribute
	       var comment string
	       if constraint.Required {
		       comment = "required"
	       } else {
		       comment = "optional"
	       }
	       if constraint.MaxLength > 0 {
		       comment += fmt.Sprintf(", max %d chars", constraint.MaxLength)
	       }
	       if constraint.Format != "" {
		       comment += fmt.Sprintf(", format: %s", constraint.Format)
	       }
	       if len(constraint.Enum) > 0 {
		       comment += fmt.Sprintf(", enum(%d values)", len(constraint.Enum))
	       }
	       content.WriteString(fmt.Sprintf("  // %s\n", comment))
	       content.WriteString(fmt.Sprintf("  attribute \"%s\" {\n", tfAttrName))
	       if sourceRequest == "manual" && constraint.ApiRef != "" {
		       content.WriteString(fmt.Sprintf("    api_ref = \"%s\"\n", constraint.ApiRef))
	       } else {
		       requestType := info.CreateRequestType
		       if sourceRequest != "" && sourceRequest != "manual" {
			       requestType = sourceRequest
		       }
		       content.WriteString(fmt.Sprintf("    api_ref = \"%s.%s\"\n", requestType, propName))
	       }
		
		// Add constraints if present
		if constraint.MaxLength > 0 {
			content.WriteString(fmt.Sprintf("    max_length = %d\n", constraint.MaxLength))
		}
		if constraint.MinLength > 0 {
			content.WriteString(fmt.Sprintf("    min_length = %d\n", constraint.MinLength))
		}
		if constraint.Pattern != "" {
			// Escape backslashes for HCL
			escapedPattern := strings.ReplaceAll(constraint.Pattern, `\`, `\\`)
			content.WriteString(fmt.Sprintf("    pattern = \"%s\"\n", escapedPattern))
		}
		if len(constraint.Enum) > 0 {
			// Write enum values as an HCL list
			content.WriteString("    valid_values = [")
			for i, val := range constraint.Enum {
				if i > 0 {
					content.WriteString(", ")
				}
				content.WriteString(fmt.Sprintf("%q", val))
			}
			content.WriteString("]\n")
		}
		
		content.WriteString("  }\n\n")
	}
	
	content.WriteString("  // Add manual customizations below with // MANUAL: comment\n")
	content.WriteString("  // Example:\n")
	content.WriteString("  // // MANUAL: custom constraint\n")
	content.WriteString("  // attribute \"display_name\" {\n")
	content.WriteString("  //   api_ref = \"CreateXxxRequest.displayName\"\n")
	content.WriteString("  //   max_length = 256\n")
	content.WriteString("  //   pattern = \"^[a-zA-Z0-9_]+$\"\n")
	content.WriteString("  //   warn_on_exceed = true\n")
	content.WriteString("  // }\n")
	content.WriteString("}\n")
	
	if err := os.WriteFile(filename, []byte(content.String()), 0644); err != nil {
		return "", err
	}
	
	if fileExists && !strings.HasSuffix(filename, ".new") {
		return "updated", nil
	}
	return "created", nil
}

func generateComparisonReport(resourceName, existingContent string, info *ResourceInfo, outputPath string) {
	reportFile := filepath.Join(outputPath, resourceName+".diff.txt")
	
	var report strings.Builder
	report.WriteString(fmt.Sprintf("Comparison Report for %s\n", resourceName))
	report.WriteString(strings.Repeat("=", 80) + "\n\n")
	report.WriteString("EXISTING FILE has manual customizations.\n")
	report.WriteString("NEW constraints from API spec:\n\n")
	
	var propNames []string
	for name := range info.Constraints {
		propNames = append(propNames, name)
	}
	sort.Strings(propNames)
	
	for _, propName := range propNames {
		constraint := info.Constraints[propName]
		tfAttrName := toSnakeCase(propName)
		
		report.WriteString(fmt.Sprintf("  %s:\n", tfAttrName))
		if constraint.Required {
			report.WriteString("    - required: true\n")
		}
		if constraint.MaxLength > 0 {
			report.WriteString(fmt.Sprintf("    - max_length: %d\n", constraint.MaxLength))
		}
		if constraint.Format != "" {
			report.WriteString(fmt.Sprintf("    - format: %s\n", constraint.Format))
		}
		if len(constraint.Enum) > 0 {
			report.WriteString(fmt.Sprintf("    - enum: %d values\n", len(constraint.Enum)))
		}
	}
	
	report.WriteString("\n" + strings.Repeat("=", 80) + "\n")
	report.WriteString("ACTION REQUIRED:\n")
	report.WriteString("1. Review the .new file with updated constraints\n")
	report.WriteString("2. Manually merge with your customizations\n")
	report.WriteString("3. Delete .new and .diff.txt files when done\n")
	
	os.WriteFile(reportFile, []byte(report.String()), 0644)
}

func generateSummaryReport(resourceMap map[string]*ResourceInfo) {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("CONSTRAINT SUMMARY")
	fmt.Println(strings.Repeat("=", 80))
	
	// Sort by resource name
	var names []string
	for name := range resourceMap {
		names = append(names, name)
	}
	sort.Strings(names)
	
	for _, name := range names {
		info := resourceMap[name]
		tfName := "fabric_" + toSnakeCase(name)
		
		fmt.Printf("\n%s (%s):\n", tfName, info.CreateRequestType)
		
		var props []string
		for prop := range info.Constraints {
			props = append(props, prop)
		}
		sort.Strings(props)
		
		for _, prop := range props {
			c := info.Constraints[prop]
			tfProp := toSnakeCase(prop)
			
			details := []string{}
			if c.Required {
				details = append(details, "required")
			}
			if c.MaxLength > 0 {
				details = append(details, fmt.Sprintf("max: %d", c.MaxLength))
			}
			if c.Format != "" {
				details = append(details, fmt.Sprintf("format: %s", c.Format))
			}
			if len(c.Enum) > 0 {
				details = append(details, fmt.Sprintf("enum(%d values)", len(c.Enum)))
			}
			
			fmt.Printf("  - %s: %s\n", tfProp, strings.Join(details, ", "))
		}
	}
}

func toSnakeCase(s string) string {
	// Handle common acronyms first
	acronyms := map[string]string{
		"SQL":     "sql",
		"KQL":     "kql",
		"ML":      "ml",
		"API":     "api",
		"ID":      "id",
		"URL":     "url",
		"URI":     "uri",
		"HTTP":    "http",
		"HTTPS":   "https",
		"GraphQL": "graphql",
	}
	
	// Check if the entire string is an acronym
	if replacement, exists := acronyms[s]; exists {
		return replacement
	}
	
	// Replace known acronyms within the string
	result := s
	for acronym, replacement := range acronyms {
		result = strings.ReplaceAll(result, acronym, strings.Title(replacement))
	}
	
	// Now convert camelCase/PascalCase to snake_case
	var snake []rune
	runes := []rune(result)
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		
		// Add underscore before uppercase letter if:
		// 1. Not the first character
		// 2. Previous character is lowercase OR next character is lowercase (handles acronyms)
		if i > 0 && r >= 'A' && r <= 'Z' {
			prevIsLower := runes[i-1] >= 'a' && runes[i-1] <= 'z'
			nextIsLower := i+1 < len(runes) && runes[i+1] >= 'a' && runes[i+1] <= 'z'
			
			if prevIsLower || nextIsLower {
				snake = append(snake, '_')
			}
		}
		
		snake = append(snake, r)
	}
	
	return strings.ToLower(string(snake))
}

