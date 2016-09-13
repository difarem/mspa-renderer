# MSPA Panel Renderer

Utility for rendering txt files from MS PaintAdventures (such as http://www.mspaintadventures.com/6/001901.txt).

## Installation

Just run `go get -u github.com/difarem/mspa-renderer` with your GOPATH set up.

## Running

Run inside the root repository directory (`$GOPATH/src/github.com/difarem/mspa-renderer`). There are
two modes available:

* Web mode: `mspa-renderer web <listen address>`. Runs the program as a web server that dynamically
	fetches the txt files from MSPA and renders them.
* Batch mode: `mspa-renderer batch <input dir> <output dir>`. Renders all txt files in the input directory
	(in the format `<input dir>/<adventure id>/<page id>.txt`) and places them as HTML files in the
	output directory.
