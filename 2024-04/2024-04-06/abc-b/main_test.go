package main

import (
	"reflect"
	"testing"
)

func Test_mostFarDot(t *testing.T) {
	type args struct {
		dots [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test",
			args: args{dots: [][]int{
				{0, 0},
				{1, 1},
			}},
			want: []int{2, 1},
		},
		{
			name: "test",
			args: args{dots: [][]int{
				{0, 0},
				{1, 0},
				{-1, 0},
			}},
			want: []int{2, 3, 2},
		},
		{
			name: "test",
			args: args{dots: [][]int{
				{0, 0},
				{2, 4},
				{5, 0},
				{3, 4},
			}},
			want: []int{3, 3, 1, 1},
		},
		{
			name: "test",
			args: args{dots: [][]int{
				{3, 2},
				{1, 6},
				{4, 5},
				{1, 3},
				{5, 5},
				{9, 8},
			}},
			want: []int{6, 6, 6, 6, 6, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostFarDot(tt.args.dots); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mostFarDot() = %v, want %v", got, tt.want)
			}
		})
	}
}
