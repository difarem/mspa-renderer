package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("view.tmpl"))

func viewHandler(w http.ResponseWriter, r *http.Request) {
	type bindings struct {
		S         string
		P         string
		PrevPage  string
		Page      *Page
		NextNames []string
	}
	var b bindings

	b.S = r.URL.Query().Get("s")
	b.P = r.URL.Query().Get("p")

	if resp, err := http.Get(fmt.Sprintf("http://www.mspaintadventures.com/%s/%s.txt", b.S, b.P)); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else {
		defer resp.Body.Close()
		b.Page = LoadPage(resp.Body)
	}

	if err := templates.ExecuteTemplate(w, "view.tmpl", b); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func runWeb(listenAddr string) {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(listenAddr, nil)
}
