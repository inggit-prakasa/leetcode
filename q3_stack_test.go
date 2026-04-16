package main

import "testing"

func Test_exclusiveTime(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		logs     []string
		expected []int
	}{
		{
			name: "example 1",
			n:    2,
			logs: []string{
				"0:start:0",
				"1:start:2",
				"1:end:5",
				"0:end:6",
			},
			expected: []int{3, 4},
		},
		{
			name: "example 2",
			n:    1,
			logs: []string{
				"0:start:0",
				"0:end:0",
			},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := exclusiveTime(tt.n, tt.logs)
			if !compareSlices(result, tt.expected) {
				t.Errorf("exclusiveTime() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
