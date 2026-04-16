package main

import "testing"

func Test_MinimumPairRemoval(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{5,2,3,1},
			expected: 2,
		},
		{
			name:     "example 2",
			nums:     []int{1,2,2},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := minimumPairRemoval(tt.nums)
			if ans != tt.expected {
				t.Errorf("minimumPairRemoval(%v) = %d; want %d", tt.nums, ans, tt.expected)
			}
		})
	}
}
