package story

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	f, _ := os.Open("testdata/gopher.json")
	s, start, err := NewStory(f)
	if err != nil {
		t.Error(err)
	}

	if start.Title != "The Little Blue Gopher" {
		t.Error("Title not The Little Blue Gopher")
	}

	n, err := s.Next("new-york")
	if err != nil {
		t.Error(err)
	}
	if n.Title != "Visiting New York" {
		t.Error(` n.Title != "Visiting New York"`)
	}

}
