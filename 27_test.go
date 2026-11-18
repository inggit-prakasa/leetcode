package main

import "testing"

func Test_removeElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: 2,
		},
		{
			name:     "example 2",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeElement(tt.nums, tt.val)
			if result != tt.expected {
				t.Errorf("removeElement(%v, %d) = %d, want %d", tt.nums, tt.val, result, tt.expected)
			}
		})
	}
}
