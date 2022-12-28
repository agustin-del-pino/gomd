package gomd

import (
	"fmt"
	"strings"
)

type Paragraph struct {
	Text string
}

func (p Paragraph) String() string {
	return p.Text
}

type Quote struct {
	Paragraph Paragraph
}

func (q Quote) String() string {
	return fmt.Sprintf("> %s", q.Paragraph)
}

type Header struct {
	Heading   int
	Paragraph Paragraph
}

func (h Header) String() string {
	return fmt.Sprintf("%s %s", strings.Repeat("#", h.Heading), h.Paragraph)
}

type ListItem struct {
	Paragraph Paragraph
	SubList   *BulletList
}

func (i *ListItem) string(_i int) string {
	if len(i.SubList.Items) == 0 {
		return fmt.Sprintf("- %s", i.Paragraph)
	} else {
		b := strings.Builder{}
		b.WriteString(fmt.Sprintf("- %s\n%s", i.Paragraph, i.SubList.string(_i)))
		return b.String()
	}
}

func (i ListItem) String() string {
	if len(i.SubList.Items) == 0 {
		return fmt.Sprintf("- %s", i.Paragraph)
	} else {
		b := strings.Builder{}
		b.WriteString(fmt.Sprintf("- %s\n%s", i.Paragraph, i.SubList.string(2)))
		return b.String()
	}
}

func NewListItem(t string) *ListItem {
	return &ListItem{
		Paragraph: Paragraph{Text: t},
		SubList: NewBulletList(),
	}
}

type BulletList struct {
	Items []*ListItem
}

func (l *BulletList) AddItem(i *ListItem) *BulletList {
	l.Items = append(l.Items, i)
	return l
}

func NewBulletList() *BulletList {
	return &BulletList{
		Items: make([]*ListItem, 0),
	}
}

func (l BulletList) string(i int) string {
	b := strings.Builder{}

	for _, _i := range l.Items {
		b.WriteString(fmt.Sprintf("%s %s\n", strings.Repeat(" ", i), _i.string(i+2)))
	}

	return b.String()
}

func (l BulletList) String() string {
	b := strings.Builder{}

	for _, i := range l.Items {
		b.WriteString(fmt.Sprintf("%s\n", i.String()))
	}

	return b.String()
}

type CodeBlock struct {
	Lang string
	Code string
}

func (c CodeBlock) String() string {
	return fmt.Sprintf("````%s\n%s\n````", c.Lang, c.Code)
}
