package main

import "testing"

func Test_minimumOperation(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "basic test case",
			nums:     []int{1, 2, 3, 4},
			expected: 3,
		},
		{
			name:     "all elements already divisible by three",
			nums:     []int{3, 6, 9},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimumOperations(tt.nums)
			if result != tt.expected {
				t.Errorf("minimumOperations(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
