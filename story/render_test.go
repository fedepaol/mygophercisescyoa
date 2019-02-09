package story

import (
	"bufio"
	"bytes"
	"flag"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var update = flag.Bool("update", false, "update .golden files")

func TestRender(t *testing.T) {
	r := newRenderer("./testdata/template.tpl")

	n := Node{
		Title: "Title",
		Story: []string{"Story", "Story1"},
		Options: []NodeOption{
			NodeOption{
				"aa",
				"bb"},
		},
	}

	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	r.render(w, "localhost:8080", n)
	w.Flush()

	golden := filepath.Join("testdata", "render.golden")
	if *update {
		ioutil.WriteFile(golden, b.Bytes(), 0644)
	}
	expected, _ := ioutil.ReadFile(golden)
	t.Error("Output different from golden file")

	if !bytes.Equal(b.Bytes(), expected) {
		t.Error("Output different from golden file")
	}

}
