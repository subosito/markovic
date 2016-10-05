package cmark_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subosito/markovic/cmark"
)

func TestMarkdownHTML(t *testing.T) {
	s := cmark.MarkdownHTML("# Hello", 0)
	assert.Equal(t, "<h1>Hello</h1>\n", s)
}
