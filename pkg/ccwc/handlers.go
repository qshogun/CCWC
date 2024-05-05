package ccwc

import (
	"fmt"
)

func HandleCommand(handler func(string) int, data string, commandName string) {
	result := handler(data)
	fmt.Printf("%s: %d\n", commandName, result)
}
