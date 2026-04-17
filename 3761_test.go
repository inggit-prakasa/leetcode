package main

import "testing"

func Test_minMirrorPairDistance(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{nums: []int{120, 21}},
			want: 1,
		},
		{
			name: "example 2",
			args: args{nums: []int{21, 120}},
			want: -1,
		},
		{
			name: "example 3",
			args: args{nums: []int{123, 456, 789}},
			want: -1,
		},
		{
			name: "example 4",
			args: args{nums: []int{12, 21, 45, 33, 54}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMirrorPairDistance(tt.args.nums); got != tt.want {
				t.Errorf("minMirrorPairDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {

	tests := []struct {
		name string
		args int
		want int
	}{
		{
			name: "example 1",
			args: 123,
			want: 321,
		},
		{
			name: "example 2",
			args: 120,
			want: 21,
		},
		{
			name: "example 3",
			args: 21,
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}

		})
	}
}
