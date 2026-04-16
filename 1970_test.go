package main

import "testing"

func Test_latestDayToCross(t *testing.T) {
	tests := []struct {
		name     string
		row      int
		col      int
		cells    [][]int
		expected int
	}{
		{
			name:     "example 1",
			row:      2,
			col:      2,
			cells:    [][]int{{1, 1}, {2, 1}, {1, 2}, {2, 2}},
			expected: 2,
		},
		{
			name:     "example 2",
			row:      2,
			col:      2,
			cells:    [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := latestDayToCross(tt.row, tt.col, tt.cells)
			if result != tt.expected {
				t.Errorf("latestDayToCross() = %v, want %v", result, tt.expected)
			}
		})
	}
}
