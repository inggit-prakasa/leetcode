package main

import (
	"testing"
)

func Test_maxDotProduct(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums1: []int{2, 1, -2, 5},
				nums2: []int{3, 0, -6},
			},
			want: 18,
		},
		{
			name: "2",
			args: args{
				nums1: []int{3, -2},
				nums2: []int{2, -6, 7},
			},
			want: 21,
		},
		{
			name: "3",
			args: args{
				nums1: []int{-1, -1},
				nums2: []int{1, 1},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDotProduct(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("maxDotProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
