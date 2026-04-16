package main

import (
	"testing"
)

func TestMaxMatrixSum(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected int64
	}{
		{
			name: "example 1",
			matrix: [][]int{
				{1, -1},
				{-1, 1},
			},
			expected: 4,
		},
		{
			name: "example 2",
			matrix: [][]int{
				{1, 2, 3},
				{-1, -2, -3},
				{1, 2, 3},
			},
			expected: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxMatrixSum(tt.matrix)
			if result != tt.expected {
				t.Errorf("maxMatrixSum() = %v, want %v", result, tt.expected)
			}
		})
	}
}
