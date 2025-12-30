package main

import "testing"

func TestNumMagicSquaresInside(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "example 1",
			grid: [][]int{
				{4, 3, 8, 4},
				{9, 5, 1, 9},
				{2, 7, 6, 2},
			},
			expected: 1,
		},
		{
			name: "example 2",
			grid: [][]int{
				{8},
			},
			expected: 0,
		},
		{
			name: "example 3",
			grid: [][]int{
				{5, 5, 5},
				{5, 5, 5},
				{5, 5, 5},
			},
			expected: 0,
		},
		{
			name: "example 4",
			grid: [][]int{
				{1, 8, 6},
				{10, 5, 0},
				{4, 2, 9},
			},
			expected: 0,
		},

		{
			name: "example 5",
			grid: [][]int{
				{2, 7, 6, 9},
				{9, 5, 1, 6},
				{4, 3, 8, 8},
				{1, 4, 10, 1},
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numMagicSquaresInside(tt.grid)
			if result != tt.expected {
				t.Errorf("numMagicSquaresInside(%v) = %d; want %d", tt.grid, result, tt.expected)
			}
		})
	}
}
