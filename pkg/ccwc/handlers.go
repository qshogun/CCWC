package ccwc

import (
	"fmt"
	"os"
)

func HandleCommand(function func(string) (int, error), filename string, command string) {
	result, err := function(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	output := fmt.Sprintf("%s : %d\n", command, result)
	fmt.Print(output)
}
