package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/difarem/mspa-renderer/mspa"
	"github.com/difarem/mspa-renderer/mspa/render"
)

func runBatch(path string, outdir string) {
	dir, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	fl, err := dir.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if !fl.IsDir() {
		log.Fatalf("'%s' is not a directory.", path)
	}
	fls, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	for i := range fls {
		if fls[i].IsDir() {
			s := fls[i].Name()
			// prepare output directory
			os.MkdirAll(filepath.Join(outdir, s), os.ModeDir|0775)

			panels := iterateAdventure(s, filepath.Join(path, s))

			for p := range panels {
				log.Printf("Rendering %s/%s...\n", s, p)

				// render pages
				var b render.Bindings
				b.S = s
				b.P = p
				b.Link = func(s, p string) string {
					return fmt.Sprintf("../%s/%s.html", s, p)
				}
				b.Panel = panels[p]

				// loads the name of the next pages
				for i := range b.Panel.Next {
					p := panels[b.Panel.Next[i]]
					b.NextPanels = append(b.NextPanels, p)
				}

				// tries to load the previous page
				if pn, err := strconv.Atoi(b.P); err == nil {
					ppid := fmt.Sprintf("%06d", pn-1)
					if _, ok := panels[ppid]; ok {
						b.PrevPanelID = ppid
					}
				}

				out, err := os.OpenFile(filepath.Join(outdir, s, p+".html"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0775)
				if err != nil {
					log.Print(err)
					continue
				}
				defer out.Close()
				err = render.RenderPanel(templates.Lookup("view.tmpl"), out, &b)
				if err != nil {
					log.Print(err)
				}
			}
		}
	}
}

func iterateAdventure(s string, path string) map[string]*mspa.Panel {
	pages := make(map[string]*mspa.Panel)

	dir, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	fls, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	for i := range fls {
		if !fls[i].IsDir() {
			ext := filepath.Ext(fls[i].Name())
			if ext == ".txt" {
				p := strings.TrimSuffix(fls[i].Name(), ext)
				f, err := os.Open(filepath.Join(path, fls[i].Name()))
				if err != nil {
					log.Print(err)
					continue
				}
				defer f.Close()
				pages[p] = mspa.LoadPanel(f)
			}
		}
	}
	return pages
}
