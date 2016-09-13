package render

import (
	"html/template"
	"io"

	"github.com/difarem/mspa-renderer/mspa"
)

// LinkFunc defines a function that transforms a panel ID to the URL where it's located.
type LinkFunc func(s, p string) string

// Bindings defines a panel along with all data necessary to render it.
type Bindings struct {
	// Adventure ID
	S string
	// Panel ID
	P string
	// ID of the previous panel
	PrevPanelID string
	// The actual panel to be rendered
	Panel *mspa.Panel
	// The panels the next arrows link to
	NextPanels []*mspa.Panel

	// Function used to transform panel IDs to URLs
	Link LinkFunc
}

// RenderPanel generates HTML from a Bindings object and a HTML template and passes it to a writer.
func RenderPanel(template *template.Template, w io.Writer, b *Bindings) error {
	return template.Execute(w, b)
}
