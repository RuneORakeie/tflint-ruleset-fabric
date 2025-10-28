package integration

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
	"text/template"

	"github.com/google/go-cmp/cmp"
)

type meta struct {
	Version string
}

type result struct {
	Issues []issue `json:"issues"`
	Errors []any   `json:"errors"`
}

type issue struct {
	Rule    any `json:"rule"`
	Message any `json:"message"`
	Range   any `json:"range"`
	Callers any `json:"callers"`
}

func TestIntegration(t *testing.T) {
	cases := []struct {
		Name    string
		Command *exec.Cmd
		Dir     string
	}{
		{
			Name:    "basic",
			Command: exec.Command("tflint", "--format", "json", "--force"),
			Dir:     "basic",
		},
		// {
		// 	Name:    "workspace_validation",
		// 	Command: exec.Command("tflint", "--format", "json", "--force"),
		// 	Dir:     "workspace_validation",
		// },
	}

	dir, _ := os.Getwd()
	defer func() { _ = os.Chdir(dir) }()

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			testDir := filepath.Join(dir, tc.Dir)
			if err := os.Chdir(testDir); err != nil {
				t.Fatal(err)
			}
			defer func() { _ = os.Chdir(dir) }()

			var stdout, stderr bytes.Buffer
			tc.Command.Stdout = &stdout
			tc.Command.Stderr = &stderr

			if err := tc.Command.Run(); err != nil {
				t.Fatalf("Failed: %s, stdout=%s stderr=%s", err, stdout.String(), stderr.String())
			}

			ret, err := readResultFile(testDir)
			if err != nil {
				t.Fatal(err)
			}

			var expected result
			if err := json.Unmarshal(ret, &expected); err != nil {
				t.Fatal(err)
			}

			var got result
			if err := json.Unmarshal(stdout.Bytes(), &got); err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(got, expected) {
				t.Fatalf("diff=%s", cmp.Diff(expected, got))
			}
		})
	}
}

func readResultFile(dir string) ([]byte, error) {
	resultFile := "result.json"
	if runtime.GOOS == "windows" {
		if _, err := os.Stat(filepath.Join(dir, "result_windows.json")); !os.IsNotExist(err) {
			resultFile = "result_windows.json"
		}
	}

	tmplFile := filepath.Join(dir, resultFile+".tmpl")
	if _, err := os.Stat(tmplFile); !os.IsNotExist(err) {
		return executeTemplate(tmplFile)
	}

	return os.ReadFile(filepath.Join(dir, resultFile))
}

func executeTemplate(file string) ([]byte, error) {
	want := new(bytes.Buffer)
	tmpl := template.Must(template.ParseFiles(file))
	if err := tmpl.Execute(want, meta{Version: "0.1.0"}); err != nil {
		return nil, err
	}
	return want.Bytes(), nil
}
