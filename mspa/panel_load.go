package mspa

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"strings"
)

// LoadPanel loads a panel from a reader
func LoadPanel(r io.Reader) *Panel {
	p := new(Panel)
	scanner := bufio.NewScanner(r)
	var content bytes.Buffer

	// hussie why
	wwwReplacer := strings.NewReplacer("www.", "cdn.")

	state := 0
	isLog := false
	for scanner.Scan() {
		line := scanner.Text()

		if line == "###" {
			if state == 4 {
				if isLog {
					content.WriteString(`</p></td></tr></table></div>`)
				}

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
				if strings.HasPrefix(line, "|") {
					logtype := strings.Title(strings.ToLower(line[1 : len(line)-2]))
					content.WriteString(`<div style="border:1px dashed gray;padding:1px;">`)
					content.WriteString(`<div><button type="button" class="button"
onmouseover="this.sv=this.style.backgroundColor; this.style.backgroundColor='#777777';"
onmouseout="if(this.sv)this.style.backgroundColor=this.sv; else this.style.backgroundColor='';"
onclick="this.parentNode.parentNode.childNodes[1].style.display = ''; this.parentNode.style.display = 'none'; return false;"
title="Click to show the text.">`)
					content.WriteString(fmt.Sprintf("Show %s", logtype))
					content.WriteString(`</button></div>`)
					content.WriteString(`<div class="spoiler" style="display:none">`)
					content.WriteString(`<div><button type="button" class="button"
onclick="this.parentNode.parentNode.parentNode.childNodes[0].style.display = ''; this.parentNode.parentNode.style.display = 'none'; return false;"
title="Click to hide the spoiler.">`)
					content.WriteString(fmt.Sprintf("Hide %s", logtype))
					content.WriteString(`</button></div>`)
					content.WriteString(`<table width="90%" border="0" cellpadding="3" cellspacing="0">
<tr>
<td width="100%" valign="top">
<p style=" font-weight: bold; font-family: courier, monospace;color:#000000">`)
					isLog = true
				} else if !isLog {
					content.WriteString("<p>" + line + "</p>")
				} else {
					content.WriteString(line + "<br/>")
				}
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
