package gomd

import "fmt"

func Italic(t string) string {
	return fmt.Sprintf("*%s*", t)
}

func Bold(t string) string {
	return fmt.Sprintf("**%s**", t)
}

func BoldItalic(t string) string {
	return fmt.Sprintf("***%s***", t)
}

func Code(t string) string {
	return fmt.Sprintf("`%s`", t)
}

func Link(t string, l string) string {
	return fmt.Sprintf("[%s](%s)", t, l)
}
