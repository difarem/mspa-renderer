package main

import (
	"html/template"
	"io"
)

var templates = template.Must(template.ParseFiles("view.tmpl"))

type LinkFunc func(s, p string) string

type Bindings struct {
	S          string
	P          string
	PrevPageID string
	Page       *Page
	NextPages  []*Page

	Link LinkFunc
}

func RenderPage(w io.Writer, b *Bindings) error {
	return templates.ExecuteTemplate(w, "view.tmpl", b)
}
