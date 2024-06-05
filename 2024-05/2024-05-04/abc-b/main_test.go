package main

import (
	"reflect"
	"testing"
)

func Test_correctTypeIndex(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test",
			args: args{
				s: "aaaa",
				t: "bbbbaaaa",
			},
			want: []int{5, 6, 7, 8},
		}, {
			name: "test",
			args: args{
				s: "atcoder",
				t: "atcoder",
			},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		}, {
			name: "test",
			args: args{
				s: "a",
				t: "b",
			},
			want: []int{},
		}, {
			name: "test",
			args: args{
				s: "a",
				t: "xa",
			},
			want: []int{},
		}, {
			name: "test",
			args: args{
				s: "a",
				t: "ax",
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := correctTypeIndex(tt.args.s, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("correctTypeIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
