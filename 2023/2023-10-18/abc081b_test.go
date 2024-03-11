package main

import "testing"

func Test_countMaxControl(t *testing.T) {
	type args struct {
		n       int
		numbers []int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{3, []int{8, 12, 40}}, 2},
		{args{4, []int{5, 6, 8, 10}}, 0},
		{args{6, []int{382253568, 723152896, 37802240, 379425024, 404894720, 471526144}}, 8},
	}
	for _, tt := range tests {
		t.Run("count", func(t *testing.T) {
			if got := countMaxControl(tt.args.n, tt.args.numbers); got != tt.want {
				t.Errorf("countMaxControl() = %v, want %v", got, tt.want)
			}
		})
	}
}
