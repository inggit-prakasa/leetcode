package main

import "testing"

func Test_maxDistance(t *testing.T) {
	type args struct {
		colors []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{colors: []int{1, 1, 1, 6, 1, 1, 1}},
			want: 3,
		},
		{
			name: "test2",
			args: args{colors: []int{1, 8, 3, 8, 3}},
			want: 4,
		},
		{
			name: "test3",
			args: args{colors: []int{0, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.args.colors); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
