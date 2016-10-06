package markovic_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subosito/markovic"
	"github.com/subosito/markovic/cmark"
)

func TestHTML(t *testing.T) {
	r := strings.NewReader("# Hello")
	s := markovic.HTML(r)
	assert.Equal(t, "<h1>Hello</h1>\n", s)
}

func TestXML(t *testing.T) {
	r := strings.NewReader("# Hello")
	s := markovic.XML(r)
	assert.Equal(t, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<!DOCTYPE document SYSTEM \"CommonMark.dtd\">\n<document xmlns=\"http://commonmark.org/xml/1.0\">\n  <heading level=\"1\">\n    <text>Hello</text>\n  </heading>\n</document>\n", s)
}

func TestMan(t *testing.T) {
	r := strings.NewReader("# Hello")
	s := markovic.Man(r, 10)
	assert.Equal(t, ".SH\nHello\n", s)
}

func TestCommonMark(t *testing.T) {
	r := strings.NewReader("# Hello")
	s := markovic.CommonMark(r, 10)
	assert.Equal(t, "# Hello\n", s)
}

func TestLatex(t *testing.T) {
	r := strings.NewReader("# Hello")
	s := markovic.Latex(r, 10)
	assert.Equal(t, "\\section{Hello}\n", s)
}

func TestVersion(t *testing.T) {
	v := cmark.Version()
	assert.Equal(t, "0.26.1", v)
}
