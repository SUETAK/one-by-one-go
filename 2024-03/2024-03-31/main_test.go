package main

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test",
			args: args{nums: []int{2, 3, 1, 1, 4}},
			want: true,
		},
		{
			name: "test2",
			args: args{nums: []int{3, 2, 1, 0, 4}},
			want: false,
		}, {
			name: "test2",
			args: args{nums: []int{4, 2, 1, 0, 4}},
			want: true,
		}, {
			name: "test2",
			args: args{nums: []int{2, 0, 2, 1, 0, 4}},
			want: false,
		}, {
			name: "test2",
			args: args{nums: []int{2, 0, 3, 1, 0, 4}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
