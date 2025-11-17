package main

import "testing"

func Test_kLengthApart(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected bool
	}{
		{
			name:     "all ones are k length apart",
			nums:     []int{1, 0, 0, 0, 1, 0, 0, 1},
			k:        2,
			expected: true,
		},
		{
			name:     "ones are too close",
			nums:     []int{1, 0, 0, 1, 0, 1},
			k:        2,
			expected: false,
		},
		{
			name:     "100 zeros between ones",
			nums:     []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			k:        100,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := kLengthApart(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("kLengthApart(%v, %d) = %v, want %v", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}

// func Test_kLengthApart_EdgeCases(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		nums     []int
// 		k        int
// 		expected bool
// 	}{
// 		{
// 			name:     "k equals array length",
// 			nums:     []int{1, 0, 0, 0, 1},
// 			k:        5,
// 			expected: true,
// 		},
// 		{
// 			name:     "k is 1 with adjacent ones",
// 			nums:     []int{1, 1},
// 			k:        1,
// 			expected: false,
// 		},
// 		{
// 			name:     "large k value",
// 			nums:     []int{1, 0, 1},
// 			k:        100,
// 			expected: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			result := kLengthApart(tt.nums, tt.k)
// 			if result != tt.expected {
// 				t.Errorf("kLengthApart(%v, %d) = %v, want %v", tt.nums, tt.k, result, tt.expected)
// 			}
// 		})
// 	}
// }
