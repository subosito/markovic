package markovic

import (
	"io"
	"io/ioutil"

	"github.com/subosito/markovic/cmark"
)

func HTML(r io.Reader) string {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return ""
	}

	return cmark.HTML(string(b), cmark.OPTION_SMART)
}
