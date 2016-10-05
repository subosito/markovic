package markovic

import (
	"io"
	"io/ioutil"

	"github.com/subosito/markovic/cmark"
)

func Parse(r io.Reader) string {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return ""
	}

	return cmark.MarkdownHTML(string(b), 0)
}
