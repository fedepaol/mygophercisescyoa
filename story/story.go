package story

import (
	"encoding/json"
	"fmt"
	"io"
)

type Node struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type Story map[string]Node

func (s Story) Next(option string) (Node, error) {
	res, ok := s[option]
	if !ok {
		return Node{}, fmt.Errorf("Option %s not found", option)
	}
	return res, nil
}

func NewStory(r io.Reader) (*Story, Node, error) {
	decoder := json.NewDecoder(r)
	res := make(Story)
	err := decoder.Decode(&res)
	start := res["intro"]
	return &res, start, err
}
