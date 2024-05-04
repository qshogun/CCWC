package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/qshogun/ccwc/internal/utils"
	"github.com/qshogun/ccwc/pkg/ccwc"
)

func main() {
	cFlag := flag.Bool("c", false, "Count bytes in a file")
	lFlag := flag.Bool("l", false, "Count lines in a file")
	wFlag := flag.Bool("w", false, "Count words in a file")
	mFlag := flag.Bool("m", false, "Count characters in a file")
	flag.Parse()

	flags := []bool{*cFlag, *lFlag, *wFlag, *mFlag}

	processFlags(cFlag, lFlag, wFlag, mFlag, flags, flag.Args())
}

func processFlags(cFlag *bool, lFlag *bool, wFlag *bool, mFlag *bool, flags []bool, args []string) {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Please provide a file name")
		os.Exit(2)
	}

	filename := flag.Args()[0]

	if *cFlag {
		handlecommand(ccwc.GetBytes, filename, "Bytes")
	}

	if *lFlag {
		handlecommand(ccwc.GetLines, filename, "Lines")
	}

	if *wFlag {
		handlecommand(ccwc.GetWords, filename, "Words")
	}

	if *mFlag {
		handlecommand(ccwc.GetCharacters, filename, "Characters")
	}

	if !utils.IsAnyTrue(flags) {
		handlecommand(ccwc.GetBytes, filename, "Bytes")
		handlecommand(ccwc.GetLines, filename, "Lines")
		handlecommand(ccwc.GetWords, filename, "Words")
		handlecommand(ccwc.GetCharacters, filename, "Characters")
	}
}

func handlecommand(function func(string) (int, error), filename string, command string) {
	result, err := function(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	output := fmt.Sprintf("%s : %d\n", command, result)
	fmt.Print(output)
}
