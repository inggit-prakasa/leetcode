package main

import "testing"

func Test_prefixesDivBy5(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []bool
	}{
		{
			name:     "example 1",
			nums:     []int{0, 1, 1},
			expected: []bool{true, false, false},
		},
		{
			name:     "example 2",
			nums:     []int{1, 1, 1},
			expected: []bool{false, false, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := prefixesDivBy5(tt.nums)
			if !equalBoolSlices(result, tt.expected) {
				t.Errorf("prefixesDivBy5(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}

func equalBoolSlices(a, b []bool) bool {
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
