package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/difarem/mspa-renderer/mspa"
	"github.com/difarem/mspa-renderer/mspa/render"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	var b render.Bindings

	b.S = r.URL.Query().Get("s")
	b.P = r.URL.Query().Get("p")
	b.Link = func(s, p string) string {
		return fmt.Sprintf("?s=%s&p=%s", s, p)
	}

	if p, err := mspa.DownloadPanel(b.S, b.P); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else {
		b.Panel = p
	}

	// loads the name of the next pages
	for i := range b.Panel.Next {
		p, _ := mspa.DownloadPanel(b.S, b.Panel.Next[i])
		b.NextPanels = append(b.NextPanels, p)
	}

	// tries to load the previous page
	if pn, err := strconv.Atoi(b.P); err == nil {
		ppid := fmt.Sprintf("%06d", pn-1)
		if _, err := mspa.DownloadPanel(b.S, ppid); err == nil {
			b.PrevPanelID = ppid
		}
	}

	if err := render.RenderPanel(templates.Lookup("view.tmpl"), w, &b); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func runWeb(listenAddr string) {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(listenAddr, nil)
}
