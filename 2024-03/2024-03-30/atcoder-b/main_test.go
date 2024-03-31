package main

import "testing"

func Test_getSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{s: "aababc"},
			want: 17,
		}, {
			name: "test",
			args: args{s: "abracadabra"},
			want: 54,
		}, {
			name: "test",
			args: args{s: "yay"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSubstrings(tt.args.s); got != tt.want {
				t.Errorf("getSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
