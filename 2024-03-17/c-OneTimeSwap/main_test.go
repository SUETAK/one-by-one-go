package main

import (
	"testing"
)

func Test_oneTimeSwap(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				s: "abc",
			},
			want: 3,
		},
		{
			name: "test1",
			args: args{
				s: "aaaaa",
			},
			want: 1,
		},
		{
			name: "test1",
			args: args{
				s: "aaabc",
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oneTimeSwap(tt.args.s); got != tt.want {
				t.Errorf("oneTimeSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
