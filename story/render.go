package story

import (
	"html/template"
	"io"
)

type renderer struct {
	t *template.Template
}

type TemplateData struct {
	Node Node
	URL  template.URL
}

func newRenderer(templateFile string) *renderer {
	t := template.Must(template.ParseFiles(templateFile))
	return &renderer{t}
}

func (r *renderer) render(w io.Writer, baseURL string, n Node) error {
	err := r.t.Execute(w, TemplateData{n, template.URL(baseURL)})
	return err
}
