package story

import (
	"html/template"
	"io"
)

type Renderer struct {
	t *template.Template
}

func NewRenderer(templateFile string) {
	t := template.Must(template.ParseFiles(templateFile))
}

func (r Renderer) render(w io.Writer, n Node) {
	err := r.t.Execute(w, n)
}
