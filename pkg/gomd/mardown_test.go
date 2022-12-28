package gomd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarkdown_String(t *testing.T) {
	md := NewMarkdown()

	s := md.Title("Test").Paragraph("test test").String()

	assert.Equal(t, "# Test\ntest test\n", s)
}
