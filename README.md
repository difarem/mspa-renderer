# MSPA Panel Renderer

A utility for rendering the internal txt files from MS Paint Adventures (such as http://www.mspaintadventures.com/6/001901.txt) to HTML.

## Installation

Run `go get -u github.com/difarem/mspa-renderer` with your GOPATH set up.

## Running

The utility must be run inside the root repository directory (`$GOPATH/src/github.com/difarem/mspa-renderer`). There are
two modes available:

* Web mode: `mspa-renderer web <listen address>`. This creates a web server that dynamically
	fetches the txt files from MSPA and renders them.
* Batch mode: `mspa-renderer batch <input dir> <output dir>`. This renders all txt files in the input directory
	(in the format `<input dir>/<adventure id>/<page id>.txt`) and places them as HTML files in the
	output directory.

## Custom HTML templates

You may modify the `view.tmpl` file to change the layout of the resulting page.
