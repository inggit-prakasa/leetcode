package main

import "testing"

func Test_minimumHammingDistance(t *testing.T) {
	type args struct {
		source       []int
		target       []int
		allowedSwaps [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				source:       []int{1, 2, 3, 4},
				target:       []int{2, 1, 4, 5},
				allowedSwaps: [][]int{{0, 1}, {2, 3}},
			},
			want: 1,
		},
		{
			name: "Test Case 2",
			args: args{
				source:       []int{1, 2, 3, 4},
				target:       []int{1, 3, 2, 4},
				allowedSwaps: [][]int{},
			},
			want: 2,
		},
		{
			name: "Test Case 3",
			args: args{
				source:       []int{5, 1, 2, 4},
				target:       []int{1, 5, 4, 2},
				allowedSwaps: [][]int{{0, 4}, {4, 2}, {1, 3}, {1, 4}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumHammingDistance(tt.args.source, tt.args.target, tt.args.allowedSwaps); got != tt.want {
				t.Errorf("minimumHammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
