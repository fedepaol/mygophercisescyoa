package main

import (
	"bufio"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gophercises/fedepaol/cyoa/story"
)

func main() {
	storyDump := flag.String("story", "story.json", "the name of the file containing the dump of the story")
	templateFile := flag.String("template", "storytemplate.tpl", "the name of the template file for rendering the story")

	flag.Parse()

	storyFile, err := os.Open(*storyDump)
	if err != nil {
		log.Fatalf("Could not open file %s", *storyDump)
	}
	b := bufio.NewReader(storyFile)
	s, start, err := story.NewStory(b)
	r := story.NewRenderer(*templateFile)

	handler := story.NewHandler(*s, start, "http://localhost:8080", *r)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
