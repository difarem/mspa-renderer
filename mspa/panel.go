package mspa

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path"
)

// Panel defines a MSPA panel.
type Panel struct {
	Title   template.HTML
	Assets  []Asset
	Content template.HTML
	Next    []string
}

// Asset defines any kind of asset a panel can use (images, flashes)
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

var ErrPanelNotFound = errors.New("Panel not found")

// DownloadPanel downloads a panel from mspaintadventures.com
func DownloadPanel(s, p string) (*Panel, error) {
	resp, err := http.Get(fmt.Sprintf("http://www.mspaintadventures.com/%s/%s.txt", s, p))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrPanelNotFound
	}
	defer resp.Body.Close()
	return LoadPanel(resp.Body), err
}
