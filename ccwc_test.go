package main

import (
	"testing"
)

func TestCCWC(t *testing.T) {
	tests := []struct {
		name     string
		function func(string) (int, error)
		file     string
		expected int
	}{
		{
			name:     "GetBytes",
			function: getbytes,
			file:     "testdata/test.txt",
			expected: 335039,
		},
		{
			name:     "GetLines",
			function: getlines,
			file:     "testdata/test.txt",
			expected: 7143,
		},
		{
			name:     "GetWords",
			function: getwords,
			file:     "testdata/test.txt",
			expected: 58164,
		},
		{
			name:     "GetCharacters",
			function: getcharacters,
			file:     "testdata/test.txt",
			expected: 332143,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := test.function(test.file)
			if err != nil {
				t.Fatal(err)
			}
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}
