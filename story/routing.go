package story

import (
	"html/template"
	"net/http"
)

func StoryHandler(s Story, template *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		next, err := s.Next(path)
		if err != nil {
			// TODO 404
		}
		template.Execute(w, next)
	}
}
