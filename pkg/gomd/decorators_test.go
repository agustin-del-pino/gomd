package gomd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItalic(t *testing.T) {
	assert.Equal(t, "*test*", Italic("test"))
}

func TestBold(t *testing.T) {
	assert.Equal(t, "**test**", Bold("test"))
}


func TestBoldItalic(t *testing.T) {
	assert.Equal(t, "***test***", BoldItalic("test"))
}

func TestCode(t *testing.T) {
	assert.Equal(t, "`test`", Code("test"))
}

func TestLink(t *testing.T) {
	assert.Equal(t, "[test](www.test.com)", Link("test", "www.test.com"))
}