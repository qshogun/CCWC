package ccwc

import (
	"fmt"
	"os"

	"github.com/qshogun/ccwc/internal/utils"
)

func ProcessFlags(flags map[string]bool, args []string) {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Please provide a file name")
		os.Exit(2)
	}

	filename := args[0]

	if flags["Bytes"] {
		HandleCommand(GetBytes, filename, "Bytes")
	}

	if flags["Lines"] {
		HandleCommand(GetLines, filename, "Lines")
	}

	if flags["Words"] {
		HandleCommand(GetWords, filename, "Words")
	}

	if flags["Characters"] {
		HandleCommand(GetCharacters, filename, "Characters")
	}

	if !utils.IsAnyTrue(flags) {
		HandleCommand(GetBytes, filename, "Bytes")
		HandleCommand(GetLines, filename, "Lines")
		HandleCommand(GetWords, filename, "Words")
		HandleCommand(GetCharacters, filename, "Characters")
	}
}
