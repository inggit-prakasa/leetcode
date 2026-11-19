package main

import "testing"

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "example 1",
			s:        "III",
			expected: 3,
		},
		{
			name:     "example 2",
			s:        "IV",
			expected: 4,
		},
		{
			name:     "example 3",
			s:        "IX",
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := romanToInt(tt.s)
			if result != tt.expected {
				t.Errorf("romanToInt(%v) = %v; want %v", tt.s, result, tt.expected)
			}
		})
	}
}
