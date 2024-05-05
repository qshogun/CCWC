package main

import (
	"flag"

	"github.com/qshogun/ccwc/pkg/ccwc"
)

func main() {
	cFlag := flag.Bool("c", false, "Count bytes in a file")
	lFlag := flag.Bool("l", false, "Count lines in a file")
	wFlag := flag.Bool("w", false, "Count words in a file")
	mFlag := flag.Bool("m", false, "Count characters in a file")
	flag.Parse()

	flags := map[string]bool{
		"Bytes":      *cFlag,
		"Lines":      *lFlag,
		"Words":      *wFlag,
		"Characters": *mFlag,
	}

	ccwc.ProcessFlags(flags, flag.Args())
}
