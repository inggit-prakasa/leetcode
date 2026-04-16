package main

import "testing"

func Test_minTimeToVisitAllPoints(t *testing.T) {
	tests := []struct {
		name     string
		points   [][]int
		expected int
	}{
		{
			name:     "example 1",
			points:   [][]int{{1, 1}, {3, 4}, {-1, 0}},
			expected: 7,
		},
		{
			name:     "example 2",
			points:   [][]int{{3, 2}, {-2, 2}},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minTimeToVisitAllPoints(tt.points)
			if result != tt.expected {
				t.Errorf("minTimeToVisitAllPoints(%v) = %d; want %d", tt.points, result, tt.expected)
			}
		})
	}
}
