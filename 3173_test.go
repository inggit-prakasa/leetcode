package main

import (
	"testing"
)

func Test_longestBalanced(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "example 1",
			s:        "abbac",
			expected: 4,
		},
		{
			name:     "example 2",
			s:        "zzabccy",
			expected: 4,
		},
		{
			name:     "example 3",
			s:        "aba",
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestBalanced(tt.s)
			if result != tt.expected {
				t.Errorf("longestBalanced(%v) = %v; want %v", tt.s, result, tt.expected)
			}
		})
	}
}
