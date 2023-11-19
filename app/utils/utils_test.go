package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	var tests = []struct {
		slice    []string
		str      string
		expected bool
	}{
		{[]string{"apple", "banana", "orange", "grape"}, "orange", true},
		{[]string{"apple", "banana", "orange", "grape"}, "kiwi", false},
		{[]string{"one", "two", "three"}, "two", true},
		{[]string{"red", "green", "blue"}, "purple", false},
		{[]string{"a", "b", "c"}, "b", true},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result := Contains(tt.slice, tt.str)

			// Check if the result matches the expected value
			assert.Equal(t, tt.expected, result)
		})
	}
}
