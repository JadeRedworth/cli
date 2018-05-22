package test

import (
	"testing"

	"github.com/fnproject/cli/common"
)

func TestCleanImageName(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"someimage:latest", "someimage"},
		{"repository/image/name:latest", "repository/image/name"},
		{"repository:port/image/name:latest", "repository:port/image/name"},
	}
	for _, c := range testCases {
		t.Run(c.input, func(t *testing.T) {
			output := common.CleanImageName(c.input)
			if output != c.expected {
				t.Fatalf("Expected '%s' but got '%s'", c.expected, output)
			}
		})
	}
}
