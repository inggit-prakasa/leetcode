package main

import "testing"

func Test_maxSumDivThree(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{3, 6, 5, 1, 8},
			expected: 18,
		},
		{
			name:     "example 2",
			nums:     []int{4},
			expected: 0,
		},
		{
			name:     "example 3",
			nums:     []int{1, 2, 3, 4, 4},
			expected: 12,
		},
		{
			name:     "additional test case",
			nums:     []int{2, 6, 2, 2, 7},
			expected: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxSumDivThree(tt.nums)
			if result != tt.expected {
				t.Errorf("maxSumDivThree(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
