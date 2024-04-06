package main

import "testing"

func Test_getKickResult(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{n: 7},
			want: "ooxooxo",
		},
		{
			name: "test",
			args: args{n: 0},
			want: "",
		},
		{
			name: "test",
			args: args{n: 1},
			want: "o",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKickResult(tt.args.n); got != tt.want {
				t.Errorf("getKickResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
