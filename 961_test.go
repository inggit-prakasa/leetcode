package main

import (
	"testing"
)

func Test_repeatedNTimes(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{1, 2, 3, 3},
			expected: 3,
		},	
		{
			name:     "example 2",
			nums:     []int{2, 1, 2, 5, 3, 2},
			expected: 2,
		},
		{
			name:     "example 3",
			nums:     []int{5, 1, 5, 2, 5, 3, 5, 4},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := repeatedNTimes(tt.nums)
			if result != tt.expected {
				t.Errorf("repeatedNTimes(%v) = %d; want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
