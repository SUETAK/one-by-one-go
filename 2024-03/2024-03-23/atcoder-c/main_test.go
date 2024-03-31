package main

import "testing"

func Test_sigma(t *testing.T) {
	type args struct {
		k    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				k:    5,
				nums: []int{1, 6, 3, 1},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sigma(tt.args.k, tt.args.nums); got != tt.want {
				t.Errorf("sigma() = %v, want %v", got, tt.want)
			}
		})
	}
}
