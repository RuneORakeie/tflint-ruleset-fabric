package project

import "fmt"

const Version string = "0.1.0"

func ReferenceLink(name string) string {
	return fmt.Sprintf("https://github.com/RuneORakeie/tflint-ruleset-fabric/blob/v%s/docs/rules/%s.md", Version, name)
}
