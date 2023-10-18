package main

import (
	"testing"
)

func Test_correctBrackets(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		args args
		want map[int]string
	}{
		// TODO: Add test cases.
		{args{1}, map[int]string{0: ""}},
		{args{2}, map[int]string{0: "()"}},
		{args{3}, map[int]string{0: ""}},
		{args{4}, map[int]string{0: "(())", 1: "()()"}},
		{args{6}, map[int]string{0: "((()))", 1: "(()())", 2: "()()()"}},
	}
	for _, tt := range tests {
		t.Run("correctBrackets", func(t *testing.T) {
			got := correctBrackets(tt.args.a)
			for _, s := range got {
				for _, s2 := range tt.want {
					if s == s2 {
						break
					}
					t.Errorf("correctBrackets() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
