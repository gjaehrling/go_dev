package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const tTmpl = `// Code generated by generator. DO NOT EDIT.

package {{$.Name}}

import (
	"context"
	"testing"
)

{{range .Receivers -}}
// -----------------------------
{{if eq . "" -}}
// Free functions
{{else -}}
// Methods of helper type
type helperType struct{}
{{end -}}
// -----------------------------

func {{.}}nonTestHelper({{$.Name}} int) {}

func {{.}}helperWithoutHelper({{$.Name}} {{$.Testing}}.{{$.UName}}) {} {{if or (eq $.Check "") (eq $.Check "begin")}}// want "test helper function should start from {{$.Name}}.Helper()"{{end}}

func {{.}}helperWithHelper({{$.Name}} {{$.Testing}}.{{$.UName}}) {
	{{$.Name}}.Helper()
}

func {{.}}helperWithEmptyStringBeforeHelper({{$.Name}} {{$.Testing}}.{{$.UName}}) {

	{{$.Name}}.Helper()
}

func {{.}}helperWithHelperAfterAssignment({{$.Name}} {{$.Testing}}.{{$.UName}}) { {{if or (eq $.Check "") (eq $.Check "begin")}}// want "test helper function should start from {{$.Name}}.Helper()"{{end}}
	_ = 0
	{{$.Name}}.Helper()
}

func {{.}}helperWithHelperAfterOtherCall({{$.Name}} {{$.Testing}}.{{$.UName}}) { {{if or (eq $.Check "") (eq $.Check "begin")}}// want "test helper function should start from {{$.Name}}.Helper()"{{end}}
	ff()
	{{$.Name}}.Helper()
}

func {{.}}helperWithHelperAfterOtherSelectionCall({{$.Name}} {{$.Testing}}.{{$.UName}}) { {{if or (eq $.Check "") (eq $.Check "begin")}}// want "test helper function should start from {{$.Name}}.Helper()"{{end}}
	{{$.Name}}.Fail()
	{{$.Name}}.Helper()
}

func {{.}}helperParamNotFirst(s string, i int, {{$.Name}} {{$.Testing}}.{{$.UName}}) { {{if or (eq $.Check "") (eq $.Check "first")}}// want "parameter {{$.TestingComment}}.{{$.UName}} should be the first or after context.Context"{{end}}
	{{$.Name}}.Helper()
}

func {{.}}helperParamSecondWithoutContext(s string, {{$.Name}} {{$.Testing}}.{{$.UName}}, i int) { {{if or (eq $.Check "") (eq $.Check "first")}}// want "parameter {{$.TestingComment}}.{{$.UName}} should be the first or after context.Context"{{end}}
	{{$.Name}}.Helper()
}

func {{.}}helperParamSecondWithContext(ctx context.Context, {{$.Name}} {{$.Testing}}.{{$.UName}}) {
	{{$.Name}}.Helper()
}

func {{.}}helperWithIncorrectName(o {{$.Testing}}.{{$.UName}}) { {{if or (eq $.Check "") (eq $.Check "name")}}// want "parameter {{$.TestingComment}}.{{$.UName}} should have name {{$.Name}}"{{end}}
	o.Helper()
}

func {{.}}helperWithAnonymousHelper({{$.Name}} {{$.Testing}}.{{$.UName}}) {
	{{$.Name}}.Helper()
	func({{$.Name}} {{$.Testing}}.{{$.UName}}) {}({{$.Name}}) {{if or (eq $.Check "") (eq $.Check "begin")}}// want "test helper function should start from {{$.Name}}.Helper()"{{end}}
}

func {{.}}helperWithNoName(_ {{$.Testing}}.{{$.UName}}) {
}

{{end -}}

func ff() {}
`

func main() {
	opts := parseOptions()

	if err := os.MkdirAll(opts.path, 0o755); err != nil {
		log.Fatalf("failed to create path %q: %v", opts.path, err)
	}

	tmpl, err := template.New(opts.name).Parse(tTmpl)
	if err != nil {
		log.Fatalf("failed to parse template: %v", err)
	}

	filePath := filepath.Join(opts.path, opts.name+".go")
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("failed to create file %q: %v", filePath, err)
	}
	if err := os.MkdirAll(opts.path, 0o600); err != nil {
		log.Fatalf("failed to create directory %q: %v", opts.path, err)
	}

	testingStr := "*testing"
	testingComment := `\\*testing`
	if opts.isInterface {
		testingStr = "testing"
		testingComment = "testing"
	}
	data := struct {
		Receivers      []string
		Name           string
		UName          string
		Check          string
		Testing        string
		TestingComment string
	}{
		Receivers:      []string{"", "(h helperType) "},
		Name:           opts.name,
		UName:          strings.ToUpper(opts.name),
		Check:          opts.check,
		Testing:        testingStr,
		TestingComment: testingComment,
	}
	if err := tmpl.Execute(f, &data); err != nil {
		log.Fatalf("failed to execute template: %v", err)
	}
}

type opts struct {
	name, path, check string
	isInterface       bool
}

func parseOptions() opts {
	var opts opts
	flag.StringVar(&opts.name, "name", "", "")
	flag.StringVar(&opts.path, "path", "", "")
	flag.StringVar(&opts.check, "check", "", "")
	flag.BoolVar(&opts.isInterface, "interface", false, "")
	flag.Parse()

	if opts.name == "" || opts.path == "" {
		log.Fatal("name and path must be specified")
	}

	if opts.check != "" {
		opts.path = filepath.Join(opts.path, opts.name+"_"+opts.check)
	} else {
		opts.path = filepath.Join(opts.path, opts.name)
	}

	return opts
}