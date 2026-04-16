package main

import (
	"testing"
)

func Test_EvalRPN(t *testing.T) {
	tests := []struct {
		name     string
		tokens   []string
		expected int
	}{
		{
			name:     "example 1",
			tokens:   []string{"2", "1", "+", "3", "*"},
			expected: 9,
		},
		{
			name:     "example 2",
			tokens:   []string{"4", "13", "5", "/", "+"},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := evalRPN(tt.tokens)
			if result != tt.expected {
				t.Errorf("evalRPN(%v) = %d; want %d", tt.tokens, result, tt.expected)
			}
		})
	}
}
