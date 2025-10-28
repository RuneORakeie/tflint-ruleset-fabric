package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

type mappingFile struct {
	Mappings []mapping `hcl:"mapping,block"`
}

type mapping struct {
	Resource   string             `hcl:"resource,label"`
	ImportPath string             `hcl:"import_path"`
	Attributes []attributeMapping `hcl:"attribute,block"`
}

type attributeMapping struct {
	Name         string   `hcl:"name,label"`
	ApiRef       string   `hcl:"api_ref"`
	MaxLength    *int     `hcl:"max_length,optional"`
	MinLength    *int     `hcl:"min_length,optional"`
	Pattern      *string  `hcl:"pattern,optional"`
	WarnOnExceed *bool    `hcl:"warn_on_exceed,optional"`
	ValidValues  []string `hcl:"valid_values,optional"`
}

// manualConstraint represents manually-added constraints in mapping files
type manualConstraint struct {
	MaxLength    *int     `hcl:"max_length,optional"`
	MinLength    *int     `hcl:"min_length,optional"`
	Pattern      *string  `hcl:"pattern,optional"`
	WarnOnExceed *bool    `hcl:"warn_on_exceed,optional"`
	ValidValues  []string `hcl:"valid_values,optional"`
}

type apiSpec struct {
	definitions map[string]interface{}
	components  map[string]interface{}
}

type attributeRef struct {
	resource  string
	block     *string
	attribute string
	value     hcl.Expression
}

func (r *attributeRef) String() string {
	if r.block != nil {
		return fmt.Sprintf("%s.%s.%s", r.resource, *r.block, r.attribute)
	}
	return fmt.Sprintf("%s.%s", r.resource, r.attribute)
}

func (r *attributeRef) RuleName() string {
	// resource already has "fabric_" prefix, don't add it again
	if r.block != nil {
		return fmt.Sprintf("%s_%s_invalid_%s", r.resource, *r.block, r.attribute)
	}
	return fmt.Sprintf("%s_invalid_%s", r.resource, r.attribute)
}

func (r *attributeRef) RuleTemplate() string {
	if r.block != nil {
		return getFullPath("rule_block.go.tmpl")
	}
	return getFullPath("rule.go.tmpl")
}

type ruleMeta struct {
	RuleName      string
	RuleNameCC    string
	ResourceType  string
	BlockType     string
	AttributeName string
	Sensitive     bool
	Max           int
	SetMax        bool
	Min           int
	SetMin        bool
	MaxLength     int
	SetMaxLength  bool
	MinLength     int
	SetMinLength  bool
	Pattern       string
	Enum          []string
	Format        string
	ReadOnly      bool
	WarnOnExceed  bool
	ReferenceURL  string
}

type providerMeta struct {
	RuleNameCCList []string
}

type ruleDocMeta struct {
	RuleName      string
	ResourceType  string
	AttributeName string
	BlockType     string
	Description   string
	Severity      string
}

type ruleDocIndexMeta struct {
	RuleNameList []string
}

var BasePath string
var RulesPath string
var DocsPath string
var SpecsPath string
var schemaConstraints map[string]map[string]int // resource.attribute -> max_length
var schemaEnums map[string]map[string][]string  // resource.attribute -> valid enum values

func getFullPath(path string) string {
	return fmt.Sprintf("%s/%s", BasePath, path)
}

func parseFlags() {
	flag.StringVar(&BasePath, "base-path", "tools/apispec-rule-gen", "Base path for generator")
	flag.StringVar(&RulesPath, "rules-path", "rules", "Output path for generated rules")
	flag.StringVar(&DocsPath, "docs-path", "docs", "Output path for generated docs")
	flag.StringVar(&SpecsPath, "specs-path", "", "Path to fabric-rest-api-specs directory")
	flag.Parse()
}

var terraformSchema provider
var generatedRuleNames []string = []string{}
var generatedRuleNameCCs []string = []string{}

func main() {
	parseFlags()

	if SpecsPath == "" {
		fmt.Println("Error: -specs-path is required")
		fmt.Println("Usage: go run ./tools/apispec-rule-gen -specs-path /path/to/fabric-rest-api-specs")
		os.Exit(1)
	}

	terraformSchema = loadProviderSchema()

	// Load constraints from schema.json
	schemaConstraints = extractSchemaConstraints(terraformSchema)

	// Load enum values from schema.json
	schemaEnums = extractSchemaEnums(terraformSchema)

	// Build a set of known Terraform resources
	knownResources := make(map[string]bool)
	for resourceType := range terraformSchema.ResourceSchemas {
		knownResources[resourceType] = true
	}

	files, err := filepath.Glob(getFullPath("mappings/*.hcl"))
	if err != nil {
		panic(err)
	}

	if len(files) == 0 {
		fmt.Println("Warning: No mapping files found in", getFullPath("mappings/"))
		fmt.Println("Please create mapping files to define resource-to-spec relationships")
		os.Exit(0)
	}

	// Check for orphaned mapping files
	orphanedMappings := checkOrphanedMappings(files, knownResources)

	// Create a set of orphaned mappings for quick lookup
	orphanedSet := make(map[string]bool)
	for _, mapping := range orphanedMappings {
		orphanedSet[mapping] = true
	}

	if len(orphanedMappings) > 0 {
		fmt.Println("\n⚠️  WARNING: Found mapping files without corresponding Terraform resources:")
		for _, mapping := range orphanedMappings {
			fmt.Printf("  - %s (no resource found in schema.json)\n", mapping)
		}
		fmt.Println("\n   These mappings will be SKIPPED during rule generation.")
		fmt.Println("   Consider reviewing the provider schema or mapping file names.\n")
	}

	mappingFiles := make([]mappingFile, 0, len(files))
	for _, file := range files {
		// Extract resource name from filename and check if it's orphaned
		baseName := filepath.Base(file)
		resourceName := strings.TrimSuffix(baseName, ".hcl")

		if orphanedSet[resourceName] {
			// Skip orphaned mapping files
			continue
		}

		parser := hclparse.NewParser()
		f, diags := parser.ParseHCLFile(file)
		if diags.HasErrors() {
			panic(diags)
		}

		var mf mappingFile
		diags = gohcl.DecodeBody(f.Body, nil, &mf)
		if diags.HasErrors() {
			panic(diags)
		}
		mappingFiles = append(mappingFiles, mf)
	}

	for _, mappingFile := range mappingFiles {
		for _, mapping := range mappingFile.Mappings {
			specPath := filepath.Join(SpecsPath, mapping.ImportPath)
			raw, err := ioutil.ReadFile(specPath)
			if err != nil {
				fmt.Printf("Warning: Could not read spec file %s: %v\n", specPath, err)
				continue
			}

			var spec map[string]interface{}
			err = json.Unmarshal(raw, &spec)
			if err != nil {
				fmt.Printf("Warning: Could not parse spec file %s: %v\n", specPath, err)
				continue
			}

			// Support both Swagger 2.0 and OpenAPI 3.0
			definitions := make(map[string]interface{})
			if defs, ok := spec["definitions"].(map[string]interface{}); ok {
				definitions = defs
			}

			components := make(map[string]interface{})
			if comps, ok := spec["components"].(map[string]interface{}); ok {
				if schemas, ok := comps["schemas"].(map[string]interface{}); ok {
					// Merge OpenAPI schemas into definitions
					for k, v := range schemas {
						definitions[k] = v
					}
				}
			}

			apiSpec := apiSpec{definitions: definitions, components: components}

			// Process each attribute mapping
			for _, attr := range mapping.Attributes {
				processAttributeMapping(apiSpec, mapping, attr)
			}
		}
	}

	sort.Strings(generatedRuleNameCCs)
	generateProviderFile(generatedRuleNameCCs)
	sort.Strings(generatedRuleNames)
	generateRulesIndexDoc(generatedRuleNames)

	fmt.Printf("\n✅ Generated %d rules\n", len(generatedRuleNames))
	fmt.Printf("Rules: %s/apispec/\n", RulesPath)
	fmt.Printf("Docs: %s/rules/\n", DocsPath)

	if len(orphanedMappings) > 0 {
		fmt.Printf("\n⚠️  %d orphaned mapping files found (see warnings above)\n", len(orphanedMappings))
	}
}

// checkOrphanedMappings finds mapping files that don't have corresponding Terraform resources
func checkOrphanedMappings(files []string, knownResources map[string]bool) []string {
	var orphaned []string

	for _, file := range files {
		// Extract resource name from filename (e.g., /path/to/fabric_workspace.hcl -> fabric_workspace)
		baseName := filepath.Base(file)
		resourceName := strings.TrimSuffix(baseName, ".hcl")

		// Check if this resource exists in schema
		if !knownResources[resourceName] {
			orphaned = append(orphaned, resourceName)
		}
	}

	sort.Strings(orphaned)
	return orphaned
}

// inferRequestObject tries to infer the request object name from the resource name
// For example: fabric_activator -> CreateActivatorRequest, CreateReflex*, etc.
func inferRequestObject(resourceName string, propertyName string, apiSpec apiSpec) string {
	// Remove fabric_ prefix
	shortName := strings.TrimPrefix(resourceName, "fabric_")

	// Try common patterns
	patterns := []string{
		"Create" + toPascalCase(shortName) + "Request",
		"Create" + toPascalCase(shortName),
		toPascalCase(shortName) + "Request",
		toPascalCase(shortName),
	}

	// Also check if there's a "Reflex" variant (like activator -> Reflex)
	if shortName == "activator" {
		patterns = append(patterns, "CreateReflexRequest", "CreateReflex", "ReflexRequest", "Reflex")
	}

	// Try each pattern
	for _, pattern := range patterns {
		if defValue, ok := apiSpec.definitions[pattern]; ok && defValue != nil {
			if defMap, ok := defValue.(map[string]interface{}); ok {
				if propsValue, ok := defMap["properties"]; ok && propsValue != nil {
					if propsMap, ok := propsValue.(map[string]interface{}); ok {
						// Check if the property exists
						if _, exists := propsMap[propertyName]; exists {
							return pattern + "." + propertyName
						}
					}
				}
			}
		}
	}

	return ""
}

// toPascalCase converts snake_case to PascalCase
func toPascalCase(s string) string {
	words := strings.Split(s, "_")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(word[:1]) + word[1:]
		}
	}
	return strings.Join(words, "")
}

func processAttributeMapping(apiSpec apiSpec, mapping mapping, attr attributeMapping) {
	fmt.Printf("Generating rule for `%s.%s`\n", mapping.Resource, attr.Name)

	// Parse the API reference (e.g., "CreateLakehouseRequest.displayName")
	parts := strings.Split(attr.ApiRef, ".")

	// If only property name is provided, try to infer the request object
	if len(parts) == 1 {
		inferredApiRef := inferRequestObject(mapping.Resource, attr.ApiRef, apiSpec)
		if inferredApiRef == "" {
			fmt.Printf("Warning: Invalid API reference '%s' for %s.%s (cannot infer request object)\n", attr.ApiRef, mapping.Resource, attr.Name)
			return
		}
		fmt.Printf("  Inferred API reference: %s\n", inferredApiRef)
		parts = strings.Split(inferredApiRef, ".")
	}

	if len(parts) < 2 {
		fmt.Printf("Warning: Invalid API reference '%s' for %s.%s\n", attr.ApiRef, mapping.Resource, attr.Name)
		return
	}

	// Get the definition from the API spec
	defValue, ok := apiSpec.definitions[parts[0]]
	if !ok || defValue == nil {
		fmt.Printf("Warning: Definition '%s' not found in API spec for %s.%s\n", parts[0], mapping.Resource, attr.Name)
		return
	}

	defMap, ok := defValue.(map[string]interface{})
	if !ok {
		fmt.Printf("Warning: Definition '%s' is not a map for %s.%s\n", parts[0], mapping.Resource, attr.Name)
		return
	}

	var definition map[string]interface{}
	if len(parts) == 2 {
		// Get property from definition
		propsValue, ok := defMap["properties"]
		if !ok || propsValue == nil {
			fmt.Printf("Warning: Definition '%s' has no properties for %s.%s\n", parts[0], mapping.Resource, attr.Name)
			return
		}
		propsMap, ok := propsValue.(map[string]interface{})
		if !ok {
			fmt.Printf("Warning: Properties in '%s' is not a map for %s.%s\n", parts[0], mapping.Resource, attr.Name)
			return
		}
		propValue, ok := propsMap[parts[1]]
		if !ok || propValue == nil {
			fmt.Printf("Warning: Property '%s' not found in '%s' for %s.%s\n", parts[1], parts[0], mapping.Resource, attr.Name)
			return
		}
		definition, ok = propValue.(map[string]interface{})
		if !ok {
			fmt.Printf("Warning: Property '%s.%s' is not a map for %s.%s\n", parts[0], parts[1], mapping.Resource, attr.Name)
			return
		}
	} else {
		definition = defMap
	}

	// Create manual constraints from attribute mapping
	manualConstraints := &manualConstraint{
		MaxLength:    attr.MaxLength,
		MinLength:    attr.MinLength,
		Pattern:      attr.Pattern,
		WarnOnExceed: attr.WarnOnExceed,
		ValidValues:  attr.ValidValues,
	}

	// Check if we have valid constraints
	if validMapping(definition, manualConstraints) {
		ref := attributeRef{resource: mapping.Resource, attribute: attr.Name}
		attrSchema := extractAttrSchema(ref, definition, manualConstraints)

		// Skip if attribute not found in Terraform provider
		if attrSchema.Type == nil {
			return
		}

		meta := generateRuleFile(mapping, ref, definition, attrSchema, manualConstraints)
		generatedRuleNames = append(generatedRuleNames, meta.RuleName)
		generatedRuleNameCCs = append(generatedRuleNameCCs, meta.RuleNameCC)
	} else {
		fmt.Printf("  Skipping %s.%s - no valid constraints (UUID-only or no constraints)\n", mapping.Resource, attr.Name)
	}
}

func validMapping(definition map[string]interface{}, manualConstraints *manualConstraint) bool {
	// Check if we have manual constraints
	if manualConstraints != nil {
		if manualConstraints.MaxLength != nil || manualConstraints.MinLength != nil ||
			manualConstraints.Pattern != nil || len(manualConstraints.ValidValues) > 0 {
			return true
		}
	}

	// Otherwise check API spec definition
	defType, ok := definition["type"].(string)
	if !ok {
		return false
	}

	switch defType {
	case "string":
		// Skip readOnly fields - these are output-only and can't be validated
		if readOnly, ok := definition["readOnly"].(bool); ok && readOnly {
			return false
		}

		if _, ok := definition["enum"]; ok {
			return true
		}
		if _, ok := definition["pattern"]; ok {
			return true
		}
		if _, ok := definition["maxLength"]; ok {
			return true
		}
		if _, ok := definition["minLength"]; ok {
			return true
		}
		if _, ok := definition["format"]; ok {
			format := definition["format"].(string)
			// Skip UUID validation - these are almost always Terraform references
			// Valid formats that we can validate
			if format == "date-time" {
				return true
			}
		}
		return false
	case "integer", "number":
		if _, ok := definition["maximum"]; ok {
			return true
		}
		if _, ok := definition["minimum"]; ok {
			return true
		}
		return false
	default:
		return false
	}
}

func extractAttrSchema(ref attributeRef, definition map[string]interface{}, manualConstraints *manualConstraint) attribute {
	resourceSchema, ok := terraformSchema.ResourceSchemas[ref.resource]
	if !ok {
		// Resource not found in Terraform provider - return empty to skip
		fmt.Printf("⚠️  Warning: resource `%s` exists in API spec but not yet supported in Terraform provider\n", ref.resource)
		return attribute{}
	}
	if ref.block != nil {
		resourceSchema, ok = resourceSchema.Block.BlockTypes[*ref.block]
		if !ok {
			fmt.Printf("⚠️  Warning: block `%s.%s` exists in API spec but not yet supported in Terraform provider\n", ref.resource, *ref.block)
			return attribute{}
		}
	}
	attrSchema, ok := resourceSchema.Block.Attributes[ref.attribute]
	if !ok {
		// Return a warning instead of panic - attribute exists in API spec but not in Terraform provider yet
		fmt.Printf("⚠️  Warning: `%s` exists in API spec but not yet supported in Terraform provider\n", ref.String())
		return attribute{} // Return empty attribute to skip rule generation
	}

	// Check if definition has a type field
	defTypeRaw, hasType := definition["type"]

	// If we have manual ValidValues and no type in definition, skip type validation
	// This happens with enum references that don't have inline type definitions
	if (!hasType || defTypeRaw == nil) && manualConstraints != nil && len(manualConstraints.ValidValues) > 0 {
		// Just return the Terraform schema without validating against API definition type
		return attrSchema
	}

	if !hasType || defTypeRaw == nil {
		fmt.Printf("⚠️  Warning: `%s` has no type in API definition, skipping\n", ref.String())
		return attribute{}
	}

	defType, ok := defTypeRaw.(string)
	if !ok {
		fmt.Printf("⚠️  Warning: `%s` has non-string type in API definition: %T, skipping\n", ref.String(), defTypeRaw)
		return attribute{}
	}
	ty, ok := attrSchema.Type.(string)
	if !ok {
		panic(fmt.Sprintf("`%s` type error in schema", ref.String()))
	}

	switch defType {
	case "string":
		if ty != "string" && ty != "number" {
			panic(fmt.Sprintf("`%s` is expected as string, but got (%s)", ref.String(), attrSchema.Type))
		}
	case "integer", "number":
		if ty != "number" && ty != "string" {
			panic(fmt.Sprintf("`%s` is expected as number, but got (%s)", ref.String(), attrSchema.Type))
		}
	}

	return attrSchema
}

func generateRuleFile(mapping mapping, ref attributeRef, definition map[string]interface{}, schema attribute, manualConstraints *manualConstraint) *ruleMeta {
	ruleName := ref.RuleName()
	var blockType string
	if ref.block != nil {
		blockType = *ref.block
	}

	meta := &ruleMeta{
		RuleName:      ruleName,
		RuleNameCC:    toCamelCase(ruleName),
		BlockType:     blockType,
		ResourceType:  mapping.Resource,
		AttributeName: ref.attribute,
		Sensitive:     schema.Sensitive,
		Max:           fetchNumber(definition, "maximum"),
		SetMax:        numberExists(definition, "maximum"),
		Min:           fetchNumber(definition, "minimum"),
		SetMin:        numberExists(definition, "minimum"),
		Pattern:       fetchString(definition, "pattern"),
		Enum:          fetchStrings(definition, "enum"),
		Format:        fetchString(definition, "format"),
		ReadOnly:      fetchBool(definition, "readOnly"),
		ReferenceURL:  fmt.Sprintf("https://github.com/microsoft/fabric-rest-api-specs/tree/main/%s", strings.TrimPrefix(mapping.ImportPath, "./")),
	}

	// Priority: manual constraints > schema.json > API spec
	// Get schema.json constraint if available
	var schemaMaxLength *int
	if resourceConstraints, ok := schemaConstraints[mapping.Resource]; ok {
		attrKey := ref.attribute
		if ref.block != nil {
			attrKey = fmt.Sprintf("%s.%s", *ref.block, ref.attribute)
		}
		if maxLen, found := resourceConstraints[attrKey]; found {
			schemaMaxLength = &maxLen
		}
	}

	// Merge constraints with priority: manual > schema.json > API spec
	if manualConstraints != nil {
		if manualConstraints.MaxLength != nil {
			meta.MaxLength = *manualConstraints.MaxLength
			meta.SetMaxLength = true
		} else if schemaMaxLength != nil {
			meta.MaxLength = *schemaMaxLength
			meta.SetMaxLength = true
		} else {
			meta.MaxLength = fetchNumber(definition, "maxLength")
			meta.SetMaxLength = numberExists(definition, "maxLength")
		}

		if manualConstraints.MinLength != nil {
			meta.MinLength = *manualConstraints.MinLength
			meta.SetMinLength = true
		} else {
			meta.MinLength = fetchNumber(definition, "minLength")
			meta.SetMinLength = numberExists(definition, "minLength")
		}

		if manualConstraints.Pattern != nil {
			meta.Pattern = *manualConstraints.Pattern
		}

		if manualConstraints.WarnOnExceed != nil {
			meta.WarnOnExceed = *manualConstraints.WarnOnExceed
		}

		if len(manualConstraints.ValidValues) > 0 {
			// Filter valid_values to only include those supported by Terraform
			// Check if schema.json has enum constraints for this attribute
			if schemaEnumMap, ok := schemaEnums[ref.resource]; ok {
				if terraformEnums, ok := schemaEnumMap[ref.attribute]; ok {
					// Create a set of Terraform-allowed values
					terraformSet := make(map[string]bool)
					for _, val := range terraformEnums {
						terraformSet[val] = true
					}

					// Filter mapping values to only include Terraform-supported ones
					var filteredEnums []string
					for _, val := range manualConstraints.ValidValues {
						if terraformSet[val] {
							filteredEnums = append(filteredEnums, val)
						}
					}

					if len(filteredEnums) > 0 {
						meta.Enum = filteredEnums
						// Log if we filtered out values
						if len(filteredEnums) < len(manualConstraints.ValidValues) {
							fmt.Printf("  ℹ️  Filtered %s.%s enum from %d to %d values (Terraform-supported only)\n",
								ref.resource, ref.attribute, len(manualConstraints.ValidValues), len(filteredEnums))
						}
					}
				} else {
					// No Terraform enum constraint, use all mapping values
					meta.Enum = manualConstraints.ValidValues
				}
			} else {
				// No schema enums for this resource, use all mapping values
				meta.Enum = manualConstraints.ValidValues
			}
		}
	} else {
		// No manual constraints, check schema.json then API spec
		if schemaMaxLength != nil {
			meta.MaxLength = *schemaMaxLength
			meta.SetMaxLength = true
		} else {
			meta.MaxLength = fetchNumber(definition, "maxLength")
			meta.SetMaxLength = numberExists(definition, "maxLength")
		}
		meta.MinLength = fetchNumber(definition, "minLength")
		meta.SetMinLength = numberExists(definition, "minLength")
	}

	if meta.Pattern != "" {
		regexp.MustCompile(meta.Pattern)
	}

	generateFile(fmt.Sprintf("%s/apispec/%s.go", RulesPath, ruleName), ref.RuleTemplate(), meta)
	generateFile(fmt.Sprintf("%s/rules/%s.md", DocsPath, ruleName), getFullPath("rule.md.tmpl"), meta)

	return meta
}

func generateProviderFile(ruleNames []string) {
	meta := &providerMeta{RuleNameCCList: ruleNames}
	generateFile(fmt.Sprintf("%s/apispec/provider.go", RulesPath), getFullPath("provider.go.tmpl"), meta)
}

func generateRulesIndexDoc(ruleNames []string) {
	meta := &ruleDocIndexMeta{RuleNameList: ruleNames}
	generateFile(fmt.Sprintf("%s/README.md", DocsPath), getFullPath("doc_README.md.tmpl"), meta)
}

func fetchNumber(definition map[string]interface{}, key string) int {
	if v, ok := definition[key]; ok {
		return int(v.(float64))
	}
	return 0
}

func numberExists(definition map[string]interface{}, key string) bool {
	_, ok := definition[key]
	return ok
}

func fetchString(definition map[string]interface{}, key string) string {
	if v, ok := definition[key]; ok {
		return v.(string)
	}
	return ""
}

func fetchBool(definition map[string]interface{}, key string) bool {
	if v, ok := definition[key]; ok {
		return v.(bool)
	}
	return false
}

func fetchStrings(definition map[string]interface{}, key string) []string {
	if raw, ok := definition[key]; ok {
		list := raw.([]interface{})
		ret := make([]string, len(list))
		for i, v := range list {
			ret[i] = v.(string)
		}
		return ret
	}
	return []string{}
}

func generateFile(fileName string, tmplName string, meta interface{}) {
	dir := filepath.Dir(fileName)
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tmpl := template.Must(template.ParseFiles(tmplName))
	err = tmpl.Execute(file, meta)
	if err != nil {
		panic(err)
	}
}

var heading = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")

func toCamelCase(str string) string {
	exceptions := map[string]string{
		"ip":        "IP",
		"sql":       "SQL",
		"vm":        "VM",
		"os":        "OS",
		"id":        "ID",
		"tls":       "TLS",
		"api":       "API",
		"uuid":      "UUID",
		"url":       "URL",
		"kql":       "KQL",
		"lakehouse": "Lakehouse",
		"workspace": "Workspace",
		"fabric":    "Fabric",
	}
	parts := strings.Split(str, "_")
	replaced := make([]string, len(parts))
	for i, s := range parts {
		conv, ok := exceptions[s]
		if ok {
			replaced[i] = conv
		} else {
			replaced[i] = s
		}
	}
	str = strings.Join(replaced, "_")

	return heading.ReplaceAllStringFunc(str, func(s string) string {
		return strings.ToUpper(strings.Replace(s, "_", "", -1))
	})
}

// extractSchemaConstraints parses schema.json to find "String length must be at most X" patterns
func extractSchemaConstraints(schema provider) map[string]map[string]int {
	constraints := make(map[string]map[string]int)
	lengthPattern := regexp.MustCompile(`String length must be at most (\d+)`)

	for resourceType, resourceSchema := range schema.ResourceSchemas {
		resourceConstraints := make(map[string]int)

		// Check top-level attributes
		for attrName, attrSchema := range resourceSchema.Block.Attributes {
			if attrSchema.Description != "" {
				if matches := lengthPattern.FindStringSubmatch(attrSchema.Description); matches != nil {
					if len(matches) > 1 {
						maxLen := 0
						fmt.Sscanf(matches[1], "%d", &maxLen)
						resourceConstraints[attrName] = maxLen
					}
				}
			}
		}

		// Check block-level attributes
		for blockName, blockSchema := range resourceSchema.Block.BlockTypes {
			for attrName, attrSchema := range blockSchema.Block.Attributes {
				if attrSchema.Description != "" {
					if matches := lengthPattern.FindStringSubmatch(attrSchema.Description); matches != nil {
						if len(matches) > 1 {
							maxLen := 0
							fmt.Sscanf(matches[1], "%d", &maxLen)
							// Use block.attribute notation
							resourceConstraints[fmt.Sprintf("%s.%s", blockName, attrName)] = maxLen
						}
					}
				}
			}
		}

		if len(resourceConstraints) > 0 {
			constraints[resourceType] = resourceConstraints
		}
	}

	return constraints
}

// extractSchemaEnums parses schema.json to find "Value must be one of : X, Y, Z" patterns
func extractSchemaEnums(schema provider) map[string]map[string][]string {
	enums := make(map[string]map[string][]string)
	// Pattern to match: "Value must be one of : `Val1`, `Val2`, `Val3`"
	enumPattern := regexp.MustCompile(`Value must be one of\s*:\s*([^.]+)`)

	for resourceType, resourceSchema := range schema.ResourceSchemas {
		resourceEnums := make(map[string][]string)

		// Check top-level attributes
		for attrName, attrSchema := range resourceSchema.Block.Attributes {
			if attrSchema.Description != "" {
				if matches := enumPattern.FindStringSubmatch(attrSchema.Description); matches != nil {
					if len(matches) > 1 {
						// Extract enum values from the matched string
						// Expected format: "`Value1`, `Value2`, `Value3`"
						enumStr := matches[1]
						// Split by comma and extract values between backticks
						valuePattern := regexp.MustCompile("`([^`]+)`")
						valueMatches := valuePattern.FindAllStringSubmatch(enumStr, -1)
						var values []string
						for _, vm := range valueMatches {
							if len(vm) > 1 {
								values = append(values, vm[1])
							}
						}
						if len(values) > 0 {
							resourceEnums[attrName] = values
						}
					}
				}
			}
		}

		if len(resourceEnums) > 0 {
			enums[resourceType] = resourceEnums
		}
	}

	return enums
}
