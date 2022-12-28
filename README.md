# Go Markdown
Go API for Markdown

# Constructors

## NewMarkdown() Markdown
Instances the implementation of Markdown

````go
md := NewMarkdown()
````

# API Reference

## Markdown API

## `Header(heading int, text string) Markdown`
Adds a Header.

````go
md.Header(3, "Caption")
````

## `Title(title string) Markdown`
Adds a Title. *Header-1*

````go
md.Title("Title")
````

## `Subtitle(string) Markdown`
Adds a Subtitle. *Header-2*

````go
md.Subtitle("Subtitle")
````

## `Paragraph(string) Markdown`
Adds a Paragraph. *Just plain text*

````go
md.Paragraph("hello world")
````

## `Codeblock(lang string, code string) Markdown`
Adds a Code Block.

````go
md.Codeblock("go", "var s string")
````

## `Quote(quote string) Markdown`
Adds a Quote.

````go
md.Quote("To Be or Not To Be")
````

## `List(*BulletList) Markdown`
Adds a Bullet List.

````go
md.List(NewBulletList().AddItem(NewListItem("test")))
````

## `Headerf(heading int, format string, a ...any) Markdown`
Adds a Header with Formating.

````go
h := "Hello"
md.Headerf(9, "%s World", h)
````

## `Titlef(format string, a ...any) Markdown`
Adds a Title with Formating.

````go
md.Titlef("%s World", Bold("Hello"))
````

## `Subtitlef(format string, a ...any) Markdown`
Adds a Subtitle with Formating.

````go
md.Subtitlef("%s World", Bold("Hello"))
````

## `Paragraphf(format string, a ...any) Markdown`
Adds a Paragraph with Formating.

````go
md.Paragraphf("%s World", Bold("Hello"))
````

## `Codeblockf(lang string, format string, a ...any) Markdown`
Adds a Code Block with Formating.

````go
v := "str"
md.Codeblockf("go", "var %s string", v)
````

## `Quotef(format string, a ...any) Markdown`
Adds a Quote with Formating.

````go
s := "that's the question"
md.Quotef("To Be or Not To Be %s", s)
````

## `String() string`
Converts the Markdown into a string representation. *It's like Marshal.*

````go
s := md.String()
````
