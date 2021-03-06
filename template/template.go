package template

import (
	"html/template"
)

//go:generate go-bindata -ignore "\\.go" -pkg template -prefix dist -o bindata.go ./dist/...
//go:generate go fmt bindata.go
//go:generate sed -i.bak "s/Html/HTML/" bindata.go
//go:generate rm bindata.go.bak

// Load initializes the template files.
func Load() *template.Template {
	return template.Must(
		template.New(
			"index.html",
		).Parse(
			string(MustAsset("index.html")),
		),
	)
}
