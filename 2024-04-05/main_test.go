package main

import "testing"

func Test_hIndex(t *testing.T) {
	type args struct {
		citations []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{citations: []int{3, 0, 6, 1, 5}},
			want: 3,
		},
		{
			name: "test",
			args: args{citations: []int{0, 1}},
			want: 1,
		},
		{
			name: "test",
			args: args{citations: []int{0}},
			want: 0,
		},
		{
			name: "test",
			args: args{citations: []int{1}},
			want: 1,
		},
		{
			name: "test",
			args: args{citations: []int{1, 3, 1}},
			want: 1,
		},
		{
			name: "test",
			args: args{citations: []int{11, 15}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex(tt.args.citations); got != tt.want {
				t.Errorf("hIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
