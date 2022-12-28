package gomd

import (
	"fmt"
	"strings"
)

type Markdown interface {
	Header(int, string) Markdown
	Title(string) Markdown
	Subtitle(string) Markdown
	Paragraph(string) Markdown
	Codeblock(string, string) Markdown
	Quote(string) Markdown
	List(*BulletList) Markdown

	Headerf(int, string, ...any) Markdown
	Titlef(string, ...any) Markdown
	Subtitlef(string, ...any) Markdown
	Paragraphf(string, ...any) Markdown
	Codeblockf(string, string, ...any) Markdown
	Quotef(string, ...any) Markdown

	String() string
}

type markdown struct {
	document []interface{}
}

func (m *markdown)	Header(h int, p string) Markdown {
	m.document = append(m.document, Header{Heading: h, Paragraph: Paragraph{Text: p}})
	return m
}
func (m *markdown)	Title(t string) Markdown {
	return m.Header(1, t)
}
func (m *markdown)	Subtitle(s string) Markdown {
	return m.Header(2, s)
}
func (m *markdown)	Paragraph(p string) Markdown {
	m.document = append(m.document, Paragraph{Text: p})
	return m
}
func (m *markdown)	Codeblock(l string, c string) Markdown {
	m.document = append(m.document, CodeBlock{Lang: l, Code: c})
	return m
}
func (m *markdown)	Quote(q string) Markdown {
	m.document = append(m.document, Quote{Paragraph: Paragraph{Text: q}})
	return m
}
func (m *markdown)	List(l *BulletList) Markdown {
	m.document = append(m.document, l)
	return m
}

func (m *markdown)	Headerf(h int, f string, a ...any) Markdown {
	return m.Header(h, fmt.Sprintf(f, a...))
}
func (m *markdown)	Titlef(f string, a ...any) Markdown {
	return m.Title(fmt.Sprintf(f, a...))
}
func (m *markdown)	Subtitlef(f string, a ...any) Markdown {
	return m.Subtitle(fmt.Sprintf(f, a...))
}
func (m *markdown)	Paragraphf(f string, a ...any) Markdown {
	return m.Paragraph(fmt.Sprintf(f, a...))
}
func (m *markdown)	Codeblockf(l string, f string, a ...any) Markdown {
	return  m.Codeblock(l, fmt.Sprintf(f, a...))
}
func (m *markdown)	Quotef(f string, a ...any) Markdown {
	return m.Quote(fmt.Sprintf(f, a...))
}

func (m *markdown)	String() string {
	b := strings.Builder{}

	for _, d := range m.document {
		b.WriteString(fmt.Sprintf("%s\n", d))
	}

	return b.String()
}

func NewMarkdown() Markdown {
	return &markdown{
		document: make([]interface{}, 0),
	}
}