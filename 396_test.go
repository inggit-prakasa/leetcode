package main

import "testing"

func Test_maxRotateFunction(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{4, 3, 2, 6},
			expected: 26,
		},
		{
			name:     "example 2",
			nums:     []int{100},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxRotateFunction(tt.nums)
			if result != tt.expected {
				t.Errorf("maxRotateFunction() = %v, want %v", result, tt.expected)
			}
		})
	}

}
