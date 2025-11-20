package main

import "testing"

func Test_intersectionSizeAtLeastTwo(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  int
	}{
		{
			name:      "test 1",
			intervals: [][]int{{1, 3}, {3, 7}, {8, 9}},
			expected:  5,
		},
		{
			name:      "test 2",
			intervals: [][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}},
			expected:  3,
		},
		{
			name:      "test 3",
			intervals: [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}},
			expected:  5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := intersectionSizeTwo(tt.intervals)
			if result != tt.expected {
				t.Errorf("intersectionSizeTwo(%v) = %d, want %d", tt.intervals, result, tt.expected)
			}
		})
	}
}
