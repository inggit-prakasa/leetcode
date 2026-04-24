package main

import "testing"

func Test_furthestDistanceFromOrigin(t *testing.T) {
	tests := []struct {
		name     string
		moves    string
		expected int
	}{
		{
			name:     "example 1",
			moves:    "L_RL__R",
			expected: 3,
		},
		{
			name:     "example 2",
			moves:    "_R__LL_",
			expected: 5,
		},
		{
			name:     "example 3",
			moves:    "_______",
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := furthestDistanceFromOrigin(tt.moves)
			if result != tt.expected {
				t.Errorf("furthestDistanceFromOrigin(%v) = %d; want %d", tt.moves, result, tt.expected)
			}
		})
	}
}
