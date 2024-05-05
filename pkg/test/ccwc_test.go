package ccwc_test

import (
	"os"
	"testing"

	"github.com/qshogun/ccwc/pkg/ccwc"
)

func TestCCWC(t *testing.T) {
	tests := []struct {
		name     string
		function func(string) int
		file     string
		expected int
	}{
		{
			name:     "GetBytes",
			function: ccwc.GetBytes,
			file:     "testdata/test.txt",
			expected: 335039,
		},
		{
			name:     "GetLines",
			function: ccwc.GetLines,
			file:     "testdata/test.txt",
			expected: 7143,
		},
		{
			name:     "GetWords",
			function: ccwc.GetWords,
			file:     "testdata/test.txt",
			expected: 58164,
		},
		{
			name:     "GetCharacters",
			function: ccwc.GetCharacters,
			file:     "testdata/test.txt",
			expected: 332143,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data, _ := os.ReadFile(test.file)
			result := test.function(string(data))
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}
