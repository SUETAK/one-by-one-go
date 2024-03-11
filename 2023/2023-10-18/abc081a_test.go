package main

import "testing"

func Test_countPutOnBallCells(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{0, 0, 0}, 0},
		{args{0, 0, 1}, 1},
		{args{0, 1, 0}, 1},
		{args{0, 1, 1}, 2},
		{args{1, 0, 0}, 1},
		{args{1, 0, 1}, 2},
		{args{1, 1, 0}, 2},
		{args{1, 1, 1}, 3},
	}
	for _, tt := range tests {
		t.Run("cellCount", func(t *testing.T) {
			if got := countPutOnBallCells(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("countPutOnBallCells() = %v, want %v", got, tt.want)
			}
		})
	}
}
