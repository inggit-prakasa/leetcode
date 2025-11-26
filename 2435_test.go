package main

import "testing"

func Test_numberOfPaths(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		k        int
		expected int
	}{
		{
			name: "basic test case",
			grid: [][]int{
				{5, 2, 4},
				{3, 0, 5},
				{0, 7, 2},
			},
			k:        5,
			expected: 1,
		},
		{
			name: "single cell grid",
			grid: [][]int{
				{0,0},
			},
			k:        5,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numberOfPaths(tt.grid, tt.k)
			if result != tt.expected {
				t.Errorf("numberOfPaths(%v, %d) = %d, want %d", tt.grid, tt.k, result, tt.expected)
			}
		})
	}

}
