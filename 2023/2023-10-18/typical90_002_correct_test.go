package main

import (
	"reflect"
	"testing"
)

func Test_answerCorrectBrackets(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		args args
		want []string
	}{
		{args{1}, []string{""}},
		{args{2}, []string{"()"}},
		{args{3}, []string{""}},
		{args{4}, []string{"(())", "()()"}},
		{args{6}, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}
	for _, tt := range tests {
		t.Run("correct", func(t *testing.T) {
			if got := answerCorrectBrackets(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("answerCorrectBrackets() = %v, want %v", got, tt.want)
			}
		})
	}
}
