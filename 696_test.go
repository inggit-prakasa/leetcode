package main

import "testing"

func TestCountBinarySubstrings(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "example 1",
			s:        "00110011",
			expected: 6,
		},
		{
			name:     "example 2",
			s:        "10101",
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := countBinarySubstrings(tt.s)
			if ans != tt.expected {
				t.Errorf("countBinarySubstrings(%v) = %d; want %d", tt.s, ans, tt.expected)
			}
		})
	}
}
