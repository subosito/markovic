package markovic_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subosito/markovic"
)

func TestParse(t *testing.T) {
	r := strings.NewReader("# Hello")
	s := markovic.Parse(r)
	assert.Equal(t, "<h1>Hello</h1>\n", s)
}
