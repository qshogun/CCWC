package ccwc

import (
	"bufio"
	"os"
)

func GetBytes(filename string) (int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	return len(data), nil
}

func GetLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lines, nil
}

func GetWords(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	words := 0
	for scanner.Scan() {
		words++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return words, nil
}

func GetCharacters(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	characters := 0
	for scanner.Scan() {
		characters++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return characters, nil
}
