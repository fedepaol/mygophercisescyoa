package story

import (
	"net/http"
)

// Handler returns a handler func that handles the story
func Handler(s Story, baseURL string, rend renderer) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		next, err := s.Next(path)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		rend.render(w, baseURL, next)
	}
}
