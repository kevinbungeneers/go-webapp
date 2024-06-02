package template

import (
	"embed"
	"html/template"
	"io"
	"io/fs"
	"strings"
)

//go:embed layout page
var htmlFiles embed.FS

type Renderable interface {
	Name() string
}

type Renderer struct {
	templates *template.Template
}

func (r Renderer) Render(page Renderable, w io.Writer) error {
	return r.templates.ExecuteTemplate(w, page.Name(), page)
}

func NewRenderer() (Renderer, error) {
	var paths []string

	err := fs.WalkDir(htmlFiles, ".", func(path string, d fs.DirEntry, err error) error {
		if strings.Contains(d.Name(), ".gohtml") {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return Renderer{}, err
	}

	t, err := template.ParseFS(htmlFiles, paths...)
	if err != nil {
		return Renderer{}, err
	}

	return Renderer{templates: t}, nil
}
