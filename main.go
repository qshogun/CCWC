package main

import (
	"flag"
	"fmt"
	"os"
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
		handlecommand(getbytes, filename, "Bytes")
	}

	if *lFlag {
		handlecommand(getlines, filename, "Lines")
	}

	if *wFlag {
		handlecommand(getwords, filename, "Words")
	}

	if *mFlag {
		handlecommand(getcharacters, filename, "Characters")
	}

	if !isAnyTrue(flags) {
		handlecommand(getbytes, filename, "Bytes")
		handlecommand(getlines, filename, "Lines")
		handlecommand(getwords, filename, "Words")
		handlecommand(getcharacters, filename, "Characters")
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
