package main

import "testing"

func Test_getSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test",
			args: args{s: "commencement"},
			want: true,
		}, {
			name: "test",
			args: args{s: "banana"},
			want: false,
		}, {
			name: "test",
			args: args{s: "s"},
			want: false,
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
