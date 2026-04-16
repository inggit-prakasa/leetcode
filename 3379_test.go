package main

import (
	"testing"
)

func Test_ConstructTransformedArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "example 1",
			nums:     []int{3, -2, 1, 1},
			expected: []int{1, 1, 1, 3},
		},
		{
			name:     "example 2",
			nums:     []int{-1, 4, -1},
			expected: []int{-1, -1, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := constructTransformedArray(tt.nums)
			if !compareSlices(result, tt.expected) {
				t.Errorf("constructTransformedArray(%v) = %v; want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
