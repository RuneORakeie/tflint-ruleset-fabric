// +build ignore

package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"go/ast"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	// Find all generated rule files
	generatedDir := "rules/apispec"
	files, err := ioutil.ReadDir(generatedDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading generated directory: %v\n", err)
		os.Exit(1)
	}

	var rules []RuleInfo
	ruleFileRegex := regexp.MustCompile(`^fabric_(.+)\.go$`)

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".go") {
			matches := ruleFileRegex.FindStringSubmatch(file.Name())
			if len(matches) > 1 {
				ruleName := matches[1]
				filePath := filepath.Join(generatedDir, file.Name())

				// Parse the file to extract constructor name
				fset := token.NewFileSet()
				f, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Warning: failed to parse %s: %v\n", file.Name(), err)
					continue
				}

				// Extract constructor and type name
				constructor, typeName := extractConstructor(f)
				if constructor == "" || typeName == "" {
					fmt.Fprintf(os.Stderr, "Warning: could not extract constructor from %s\n", file.Name())
					continue
				}

				rules = append(rules, RuleInfo{
					Name:        ruleName,
					Type:        typeName,
					Constructor: constructor,
				})
			}
		}
	}

	if len(rules) == 0 {
		fmt.Fprintf(os.Stderr, "No generated rules found\n")
		os.Exit(1)
	}

	fmt.Printf("Found %d generated rules\n", len(rules))

       // Read the generated_rules_test.go file
       testFile := "rules/generated_rules_test.go"
       testContent, err := ioutil.ReadFile(testFile)
       if err != nil {
	       fmt.Fprintf(os.Stderr, "Error reading test file: %v\n", err)
	       os.Exit(1)
       }
       // Find the start and end of the generatedRuleConstructors block
       startMarker := "generatedRuleConstructors := []GeneratedRuleInfo{"
       endMarker := "// End generated rules list"
       startIdx := strings.Index(string(testContent), startMarker)
       endIdx := strings.Index(string(testContent), endMarker)
       if startIdx == -1 || endIdx == -1 {
	       fmt.Fprintf(os.Stderr, "Could not find rule registration block in test file\n")
	       os.Exit(1)
       }
       // Build the new rule registration block
       var ruleBlock strings.Builder
       ruleBlock.WriteString(startMarker + "\n")
       for _, rule := range rules {
	       ruleBlock.WriteString(fmt.Sprintf("\t{\n"))
	       ruleBlock.WriteString(fmt.Sprintf("\t\tName: \"fabric_%s\",\n", rule.Name))
	       ruleBlock.WriteString(fmt.Sprintf("\t\tType: \"%s\",\n", rule.Type))
	       ruleBlock.WriteString(fmt.Sprintf("\t\tConstructor: func() interface{ Check(tflint.Runner) error } {\n"))
	       ruleBlock.WriteString(fmt.Sprintf("\t\t\treturn %s()\n", rule.Constructor))
	       ruleBlock.WriteString(fmt.Sprintf("\t\t},\n"))
	       ruleBlock.WriteString(fmt.Sprintf("\t},\n"))
       }
       ruleBlock.WriteString("// End generated rules list\n")
       // Replace the old block with the new one
       newTestContent := string(testContent[:startIdx]) + ruleBlock.String() + string(testContent[endIdx+len(endMarker):])
       // Write back to the test file
       err = ioutil.WriteFile(testFile, []byte(newTestContent), 0644)
       if err != nil {
	       fmt.Fprintf(os.Stderr, "Error writing updated test file: %v\n", err)
	       os.Exit(1)
       }
       fmt.Println("generated_rules_test.go updated with latest rule registrations.")
}

type RuleInfo struct {
	Name        string
	Type        string
	Constructor string
}

func extractConstructor(f *ast.File) (string, string) {
	for _, decl := range f.Decls {
	       if d, ok := decl.(*ast.FuncDecl); ok && d.Name.Name != "" {
		       // Look for New* functions
		       if strings.HasPrefix(d.Name.Name, "New") {
			       constructor := d.Name.Name
			       // Extract type from return type
			       if d.Type.Results != nil && len(d.Type.Results.List) > 0 {
				       field := d.Type.Results.List[0]
				       if star, ok := field.Type.(*ast.StarExpr); ok {
					       if ident, ok := star.X.(*ast.Ident); ok {
						       return constructor, ident.Name
					       }
				       }
			       }
		       }
	       }
	}
	return "", ""
}
