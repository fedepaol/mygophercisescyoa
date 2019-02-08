package story

import (
	"encoding/json"
	"fmt"
	"io"
)

type NodeOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

// Node Represents a story node
type Node struct {
	Title   string       `json:"title"`
	Story   []string     `json:"story"`
	Options []NodeOption `json:"options"`
}

// Story is the story object
type Story map[string]Node

// Next returns the next node corresponding to the given option
func (s Story) Next(option string) (Node, error) {
	res, ok := s[option]
	if !ok {
		return Node{}, fmt.Errorf("Option %s not found", option)
	}
	return res, nil
}

// NewStory creates a new story and returns the starting node
func NewStory(r io.Reader) (*Story, Node, error) {
	decoder := json.NewDecoder(r)
	res := make(Story)
	err := decoder.Decode(&res)
	start := res["intro"]
	return &res, start, err
}
