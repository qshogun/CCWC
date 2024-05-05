package ccwc

import (
	"strings"
)

func GetBytes(data string) int {
	return len(data)
}

func GetLines(data string) int {
	return len(strings.Split(data, "\n"))
}

func GetWords(data string) int {
	return len(strings.Fields(data))
}

func GetCharacters(data string) int {
	return len([]rune(data))
}
