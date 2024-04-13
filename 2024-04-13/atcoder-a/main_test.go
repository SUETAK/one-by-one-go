package main

import "testing"

func Test_getWinnersPoint(t *testing.T) {
	type args struct {
		n  int
		ps []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				n:  4,
				ps: []int{1, -2, -1},
			},
			want: 2,
		}, {
			name: "test",
			args: args{
				n:  3,
				ps: []int{0, 0},
			},
			want: 0,
		}, {
			name: "test",
			args: args{
				n:  3,
				ps: []int{-5, -2},
			},
			want: 7,
		}, {
			name: "test",
			args: args{
				n:  6,
				ps: []int{10, 20, 30, 40, 50},
			},
			want: -150,
		}, {
			name: "test",
			args: args{
				n:  4,
				ps: []int{-7, 1, 0},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWinnersPoint(tt.args.n, tt.args.ps); got != tt.want {
				t.Errorf("getWinnersPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
