package main

import "testing"

func Test_hasAllCodes(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected bool
	}{
		{
			name:     "example 1",
			s:        "00110110",
			k:        2,
			expected: true,
		},
		{
			name:     "example 2",
			s:        "0110",
			k:        1,
			expected: true,
		},
		{
			name:     "example 3",
			s:        "0110",
			k:        2,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasAllCodes(tt.s, tt.k)
			if result != tt.expected {
				t.Errorf("hasAllCodes() = %v, want %v", result, tt.expected)
			}
		})
	}
}
