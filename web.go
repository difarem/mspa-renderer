package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var ErrPageNotFound = errors.New("Page not found")

func downloadPage(s, p string) (*Page, error) {
	resp, err := http.Get(fmt.Sprintf("http://www.mspaintadventures.com/%s/%s.txt", s, p))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrPageNotFound
	}
	defer resp.Body.Close()
	return LoadPage(resp.Body), err
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	var b Bindings

	b.S = r.URL.Query().Get("s")
	b.P = r.URL.Query().Get("p")
	b.Link = func(s, p string) string {
		return fmt.Sprintf("?s=%s&p=%s", s, p)
	}

	if p, err := downloadPage(b.S, b.P); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else {
		b.Page = p
	}

	// loads the name of the next pages
	for i := range b.Page.Next {
		p, _ := downloadPage(b.S, b.Page.Next[i])
		b.NextPages = append(b.NextPages, p)
	}

	// tries to load the previous page
	if pn, err := strconv.Atoi(b.P); err == nil {
		ppid := fmt.Sprintf("%06d", pn-1)
		if _, err := downloadPage(b.S, ppid); err == nil {
			b.PrevPageID = ppid
		}
	}

	if err := RenderPage(w, &b); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func runWeb(listenAddr string) {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(listenAddr, nil)
}
