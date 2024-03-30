package main

import (
	"reflect"
	"testing"
)

func Test_divideKValue(t *testing.T) {
	type args struct {
		k  int
		ps []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test",
			args: args{
				k:  10,
				ps: []int{50, 51, 54, 60, 65},
			},
			want: []string{"5", "6"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divideKValue(tt.args.k, tt.args.ps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divideKValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
