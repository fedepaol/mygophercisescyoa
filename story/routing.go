package story

import (
	"net/http"
)

// Handler returns a handler func that handles the story
func NewHandler(s Story, start Node, baseURL string, rend Renderer) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]
		next, err := s.Next(path)
		if err == ErrNodeNotFound {
			rend.render(w, baseURL, start)
			return
		}
		rend.render(w, baseURL, next)
	}
}
