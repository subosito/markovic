package cmark_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subosito/markovic/cmark"
)

func TestHTML(t *testing.T) {
	s := cmark.HTML("# Hello", cmark.OPTION_DEFAULT)
	assert.Equal(t, "<h1>Hello</h1>\n", s)
}

func TestXML(t *testing.T) {
	s := cmark.XML("# Hello", cmark.OPTION_DEFAULT)
	assert.Equal(t, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<!DOCTYPE document SYSTEM \"CommonMark.dtd\">\n<document xmlns=\"http://commonmark.org/xml/1.0\">\n  <heading level=\"1\">\n    <text>Hello</text>\n  </heading>\n</document>\n", s)
}

func TestMan(t *testing.T) {
	s := cmark.Man("# Hello", cmark.OPTION_DEFAULT, 10)
	assert.Equal(t, ".SH\nHello\n", s)
}

func TestCommonMark(t *testing.T) {
	s := cmark.CommonMark("# Hello", cmark.OPTION_DEFAULT, 10)
	assert.Equal(t, "# Hello\n", s)
}

func TestLatex(t *testing.T) {
	s := cmark.Latex("# Hello", cmark.OPTION_DEFAULT, 10)
	assert.Equal(t, "\\section{Hello}\n", s)
}

func TestVersion(t *testing.T) {
	v := cmark.Version()
	assert.Equal(t, "0.26.1", v)
}

func TestOptions(t *testing.T) {
	table := []struct {
		opts   cmark.Option
		input  string
		output string
	}{
		{cmark.OPTION_DEFAULT, "# Hello", "<h1>Hello</h1>\n"},
		{cmark.OPTION_SOURCEPOS, "# Hello", "<h1 data-sourcepos=\"1:1-1:7\">Hello</h1>\n"},

		{cmark.OPTION_DEFAULT, "Hello\nWorld!", "<p>Hello\nWorld!</p>\n"},
		{cmark.OPTION_HARDBREAKS, "Hello\nWorld!", "<p>Hello<br />\nWorld!</p>\n"},

		{cmark.OPTION_DEFAULT, `Hello <script>alert('hello');</script>`, "<p>Hello <script>alert('hello');</script></p>\n"},
		{cmark.OPTION_SAFE, `Hello <script>alert('hello');</script>`, "<p>Hello <!-- raw HTML omitted -->alert('hello');<!-- raw HTML omitted --></p>\n"},

		{cmark.OPTION_DEFAULT, "Hello\nWorld!", "<p>Hello\nWorld!</p>\n"},
		{cmark.OPTION_NOBREAKS, "Hello\nWorld!", "<p>Hello World!</p>\n"},

		{cmark.OPTION_DEFAULT, "# Hello--World", "<h1>Hello--World</h1>\n"},
		{cmark.OPTION_SMART, "# Hello--World", "<h1>Hello–World</h1>\n"},
	}

	for _, data := range table {
		s := cmark.HTML(data.input, data.opts)
		assert.Equal(t, data.output, s)
	}
}
