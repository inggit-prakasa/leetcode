package main

import "testing"

func Test_ClosestEqualElement(t *testing.T) {
	tests := []struct {
		name     string
		queries  []int
		arr      []int
		expected []int
	}{
		{
			name:     "example 1",
			queries:  []int{0, 3, 5},
			arr:      []int{1, 3, 1, 4, 1, 3, 2},
			expected: []int{2, -1, 3},
		},
		{
			name:     "example 2",
			queries:  []int{0, 1, 2, 3},
			arr:      []int{1, 2, 3, 4},
			expected: []int{-1, -1, -1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := solveQueries(tt.arr, tt.queries)
			if !compareSlices(result, tt.expected) {
				t.Errorf("ClosestEqualElement(%v, %v) = %v; want %v", tt.queries, tt.arr, result, tt.expected)
			}
		})
	}
}
