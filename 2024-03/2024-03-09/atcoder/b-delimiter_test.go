package atcoder

import (
	"reflect"
	"testing"
)

func Test_delimiter(t *testing.T) {
	type args struct {
		ps []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{ps: []int{3, 2, 1, 0}},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "test2",
			args: args{ps: []int{123, 456, 789, 987, 654, 321, 0}},
			want: []int{0, 321, 654, 987, 789, 456, 123},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := delimiter(tt.args.ps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("delimiter() = %v, want %v", got, tt.want)
			}
		})
	}
}
