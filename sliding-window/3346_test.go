package main

import "testing"

func Test_minimumSumSubarray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		l        int
		r        int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{3, -2, 1, 4},
			l:        2,
			r:        3,
			expected: 1,
		},
		{
			name:     "example 2",
			nums:     []int{-2, 2, -3, 1},
			l:        2,
			r:        3,
			expected: -1,
		},
		{
			name:     "example 3",
			nums:     []int{1, 2, 3, 4},
			l:        2,
			r:        4,
			expected: 3,
		},
		{
			name:     "example 4",
			nums:     []int{4, -10},
			l:        1,
			r:        1,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimumSumSubarray(tt.nums, tt.l, tt.r)
			if result != tt.expected {
				t.Errorf("minimumSumSubarray() = %d, want %d", result, tt.expected)
			}
		})
	}
}
