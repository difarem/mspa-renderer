package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		printUsage()
		return
	}

	if os.Args[1] == "batch" {
		if len(os.Args) < 4 {
			printUsage()
			return
		}
		runBatch(os.Args[2], os.Args[3])
	} else if os.Args[1] == "web" {
		runWeb(os.Args[2])
	} else {
		printUsage()
	}
}

func printUsage() {
	fmt.Printf("Usage:\nmspa-renderer batch [txt file directory] [output directory]\nmspa-renderer web [listen address]\n")
	fmt.Printf("When using batch mode, the txt files must be in the following format: [supplied directory]/[adventure id]/[panel id].txt\n")
}
