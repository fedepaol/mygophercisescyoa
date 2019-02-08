package story

import (
	"fmt"
	"html/template"
	"io"
)

type renderer struct {
	t *template.Template
}

type TemplateData struct {
	Node Node
	URL  string
}

func newRenderer(templateFile string) *renderer {
	t := template.Must(template.ParseFiles(templateFile))
	return &renderer{t}
}

func (r *renderer) render(w io.Writer, baseURL string, n Node) error {
	fmt.Printf(baseURL)
	err := r.t.Execute(w, TemplateData{n, baseURL})
	return err
}
