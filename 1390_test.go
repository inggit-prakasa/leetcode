package main

import (
	"testing"
)

func Test_sumFourDivisors(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{21, 4, 7},
			expected: 32,
		},
		{
			name:     "example 2",
			nums:     []int{21, 21},
			expected: 64,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sumFourDivisors(tt.nums)
			if result != tt.expected {
				t.Errorf("sumFourDivisors(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
