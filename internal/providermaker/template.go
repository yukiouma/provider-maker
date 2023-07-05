package providermaker

import (
	"log"
	"text/template"
)

const providerTemplate = `// Code generated by provider-maker. DO NOT EDIT.

package {{.Pkg}}

import (
	{{range $index, $import := .Imports}}
	{{$import}}{{end}}
)

var Providers = wire.NewSet({{range $index, $f := .Funcs}}
	{{$f}},{{end}}
)
`

var (
	tmpl *template.Template
	err  error
)

func init() {
	tmpl, err = template.New("provider").Parse(providerTemplate)
	if err != nil {
		log.Fatal(err)
	}
}