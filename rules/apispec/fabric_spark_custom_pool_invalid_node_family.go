package apispec

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type FabricSparkCustomPoolInvalidNodeFamily struct{ tflint.DefaultRule }

func NewFabricSparkCustomPoolInvalidNodeFamily() *FabricSparkCustomPoolInvalidNodeFamily { return &FabricSparkCustomPoolInvalidNodeFamily{} }

func (r *FabricSparkCustomPoolInvalidNodeFamily) Name() string    { return "fabric_spark_custom_pool_invalid_node_family" }
func (r *FabricSparkCustomPoolInvalidNodeFamily) Enabled() bool   { return true }
func (r *FabricSparkCustomPoolInvalidNodeFamily) Severity() string{ return tflint.ERROR }
func (r *FabricSparkCustomPoolInvalidNodeFamily) Link() string    { return "https://github.com/microsoft/fabric-rest-api-specs/tree/main/spark/definitions.json" }

func (r *FabricSparkCustomPoolInvalidNodeFamily) Check(runner tflint.Runner) error {
	resourceType := "fabric_spark_custom_pool"
	blockType    := ""     // empty string when not a nested block
	attrName     := "node_family"

	// Constraints (presence controlled by Set* flags)
	hasMinLen := false
	minLen    := 0
	hasMaxLen := false
	maxLen    := 0

	pattern   := ""
	hasRegex  := len(pattern) > 0
	var re *regexp.Regexp
	if hasRegex {
		re = regexp.MustCompile(pattern)
	}

	enum := []string{"MemoryOptimized", }
	hasEnum := len(enum) > 0

	// NOTE: .Format (uuid, uri, date-time) and .WarnOnExceed are available if you later add format-specific checks

	return helper.ForEachResource(runner, resourceType, func(res *helper.Resource) error {
		var attr *helper.Attribute

		if blockType != "" {
			blk := res.GetBlock(blockType)
			if blk == nil {
				return nil
			}
			attr = blk.GetAttribute(attrName)
		} else {
			attr = res.GetAttribute(attrName)
		}

		if attr == nil {
			return nil
		}

		// We treat values as strings for length/pattern/enum checks
		v, err := attr.ValueAsString()
		if err != nil {
			// Non-string types are typically guarded by provider schema; skip.
			return nil
		}

		// length checks
		if hasMaxLen && len(v) > maxLen {
			msg := fmt.Sprintf("%s exceeds max length %d", attrName, maxLen)
			return runner.EmitIssue(r, msg, attr.Expr.Range())
		}
		if hasMinLen && len(v) < minLen {
			msg := fmt.Sprintf("%s shorter than min length %d", attrName, minLen)
			return runner.EmitIssue(r, msg, attr.Expr.Range())
		}

		// enum
		if hasEnum {
			ok := false
			for _, ev := range enum {
				if v == ev {
					ok = true
					break
				}
			}
			if !ok {
				msg := fmt.Sprintf("%s must be one of: %s", attrName, strings.Join(enum, ", "))
				return runner.EmitIssue(r, msg, attr.Expr.Range())
			}
		}

		// regex
		if hasRegex && !re.MatchString(v) {
			msg := fmt.Sprintf("%s must match pattern %q", attrName, pattern)
			return runner.EmitIssue(r, msg, attr.Expr.Range())
		}

		return nil
	})
}