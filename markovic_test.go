package markovic_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subosito/markovic"
)

func TestHTML(t *testing.T) {
	r := strings.NewReader("# Hello")
	s := markovic.HTML(r)
	assert.Equal(t, "<h1>Hello</h1>\n", s)
}
