package story

import (
	"html/template"
	"io"
)

type Renderer struct {
	t *template.Template
}

type TemplateData struct {
	Node Node
	URL  template.URL
}

// NewRenderer returns a renderer from the given file
func NewRenderer(templateFile string) *Renderer {
	t := template.Must(template.ParseFiles(templateFile))
	return &Renderer{t}
}

func (r *Renderer) render(w io.Writer, baseURL string, n Node) error {
	err := r.t.Execute(w, TemplateData{n, template.URL(baseURL)})
	return err
}
