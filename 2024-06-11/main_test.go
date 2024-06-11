package main

import (
	"testing"
)

func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		s       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test",
			args: args{
				pattern: "abba",
				s:       "dog cat cat dog",
			},
			want: true,
		}, {
			name: "test",
			args: args{
				pattern: "abba",
				s:       "dog cat cat fish",
			},
			want: false,
		}, {
			name: "test",
			args: args{
				pattern: "ab",
				s:       "dog dog",
			},
			want: false,
		}, {
			name: "test",
			args: args{
				pattern: "a",
				s:       "dog dog",
			},
			want: false,
		}, {
			name: "test",
			args: args{
				pattern: "a",
				s:       "dog cat cat fish",
			},
			want: false,
		}, {
			name: "test",
			args: args{
				pattern: "ab",
				s:       "dog",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.s); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
