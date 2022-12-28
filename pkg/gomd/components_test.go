package gomd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParagraph(t *testing.T) {
	p := Paragraph{Text: "test"}

	assert.Equal(t, "test", p.String())
}

func TestQuote(t *testing.T) {
	q := Quote{Paragraph: Paragraph{Text: "test"}}

	assert.Equal(t, "> test", q.String())
}

func TestHeader(t *testing.T) {
	h := Header{Heading: 4, Paragraph: Paragraph{Text: "test"}}

	assert.Equal(t, "#### test", h.String())
}
func TestListItem(t *testing.T) {
	i := NewListItem("test")

	assert.Equal(t, "- test", i.String())
}
func TestBulletList(t *testing.T) {
	l := NewBulletList().AddItem(NewListItem("test")).AddItem(NewListItem("test 2"))

	assert.Equal(t, "- test\n- test 2\n", l.String())
}
func TestCodeBlock(t *testing.T) {
	c := CodeBlock{Lang: "go", Code: "var s string"}

	assert.Equal(t, "````go\nvar s string\n````", c.String())
}
