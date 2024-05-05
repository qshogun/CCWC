package ccwc

import (
	"fmt"
	"io"
	"os"

	"github.com/qshogun/ccwc/internal/utils"
)

func ProcessFlags(flags map[string]bool, args []string) {
	var data string
	var err error

	if len(args) == 0 {
		data, err = ReadFromStdin()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
	} else {
		fileData, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
		data = string(fileData)
	}

	if flags["Bytes"] {
		HandleCommand(GetBytes, data, "Bytes")
	}

	if flags["Lines"] {
		HandleCommand(GetLines, data, "Lines")
	}

	if flags["Words"] {
		HandleCommand(GetWords, data, "Words")
	}

	if flags["Characters"] {
		HandleCommand(GetCharacters, data, "Characters")
	}

	if !utils.IsAnyTrue(flags) {
		HandleCommand(GetBytes, data, "Bytes")
		HandleCommand(GetLines, data, "Lines")
		HandleCommand(GetWords, data, "Words")
		HandleCommand(GetCharacters, data, "Characters")
	}
}

func ReadFromStdin() (string, error) {
	stdin, err := io.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}
	str := string(stdin)
	return str, nil
}
