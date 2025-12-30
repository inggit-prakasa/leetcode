package main

import "testing"

func Test_specialTriplets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{6, 3, 6},
			want: 1,
		},
		{
			name: "example 2",
			nums: []int{0, 1, 0, 0},
			want: 1,
		},
		{
			name: "example 3",
			nums: []int{8, 4, 2, 8, 4},
			want: 2,
		},
		{
			name: "additional test case",
			nums: []int{3, 3, 3, 3},
			want: 0,
		},
		{
			name: "additional test case 2",
			nums: []int{84, 2, 93, 1, 2, 2, 26},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := specialTriplets(tt.nums)
			if got != tt.want {
				t.Errorf("specialTriplets(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}
