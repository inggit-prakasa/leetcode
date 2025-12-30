package main

import "testing"

func Test_minSubarray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		p        int
		expected int
	}{
		{
			name:     "example test case 1",
			nums:     []int{3, 1, 4, 2},
			p:        6,
			expected: 1,
		},
		{
			name:     "example test case 2",
			nums:     []int{6, 3, 5, 2},
			p:        9,
			expected: 2,
		},
		{
			name:     "example test case 3",
			nums:     []int{1, 2, 3},
			p:        3,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minSubarray(tt.nums, tt.p)
			if result != tt.expected {
				t.Errorf("minSubarray(%v, %d) = %d, want %d", tt.nums, tt.p, result, tt.expected)
			}
		})
	}
}