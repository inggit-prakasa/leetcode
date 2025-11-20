package main

import "testing"

func Test_findFinalValue(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		original int
		expected int
	}{
		{
			name:     "basic test case",
			nums:     []int{5, 3, 6, 1, 12},
			original: 3,
			expected: 24,
		},
		{
			name:     "no doubles",
			nums:     []int{2, 7, 9},
			original: 4,
			expected: 4,
		},
		{
			name:     "multiple doubles",
			nums:     []int{1, 2, 4, 8},
			original: 1,
			expected: 16,
		},
		{
			name:     "random order",
			nums:     []int{8, 19, 4, 2, 15, 3},
			original: 2,
			expected: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findFinalValue(tt.nums, tt.original)
			if result != tt.expected {
				t.Errorf("findFinalValue(%v, %d) = %d, want %d", tt.nums, tt.original, result, tt.expected)
			}
		})
	}
}
