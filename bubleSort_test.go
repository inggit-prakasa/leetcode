package main

import (
	"testing"
)


func Test_bubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected []int
	}{
		{
			name:     "already sorted array",
			arr:      []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "reverse sorted array",
			arr:      []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "random unsorted array",
			arr:      []int{3, 1, 4, 5, 2},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "array with duplicates",
			arr:      []int{4, 2, 5, 2, 3, 4, 1},
			expected: []int{1, 2, 2, 3, 4, 4, 5},
		},
		{
			name:     "empty array",
			arr:      []int{},
			expected: []int{},
		},
		{
			name:     "single element array",
			arr:      []int{42},
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := bubbleSort(tt.arr)
			if !equalSlices(result, tt.expected) {
				t.Errorf("bubbleSort(%v) = %v, want %v", tt.arr, result, tt.expected)
			}
		})
	}
}