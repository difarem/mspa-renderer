package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"path"
	"strings"
)

type Page struct {
	Title   template.HTML
	Assets  []Asset
	Content template.HTML
	Next    []string
}

type Asset interface {
	HTML() template.HTML
}

type Image string

func (i Image) HTML() template.HTML {
	return template.HTML(fmt.Sprintf(`<img src="%s">`, i))
}

type Flash string

func (f Flash) HTML() template.HTML {
	_, file := path.Split(string(f))

	return template.HTML(fmt.Sprintf(`<embed src="%s/%s.swf" quality="high" play="true" loop="true scale="showall" `+
		`wmode="window" devicefont="false" bgcolor="#ffffff" menu="true" allowfullscreen="false" allowscriptaccess="always" `+
		`salign="" type="application/x-shockwave-flash" width="650" align="middle" height="450"></embed>`, f, file))
}

func LoadPage(r io.Reader) *Page {
	p := new(Page)
	scanner := bufio.NewScanner(r)
	var content bytes.Buffer

	// hussie why
	wwwReplacer := strings.NewReplacer("www.", "cdn.")

	state := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "###" {
			if state == 4 {
				p.Content = template.HTML(content.String())
			}

			state++
			continue
		}

		switch state {
		case 0:
			p.Title = template.HTML(line)
		case 3:
			line = wwwReplacer.Replace(line)
			if strings.HasPrefix(line, "F|") {
				p.Assets = append(p.Assets, Flash(strings.TrimPrefix(line, "F|")))
			} else {
				p.Assets = append(p.Assets, Image(line))
			}
		case 4:
			if line != "" {
				content.WriteString("<p>" + line + "</p>")
			}
		case 5:
			if line == "X" {
				break
			}
			p.Next = append(p.Next, line)
		}
	}

	return p
}
