package atcoder

import (
	"reflect"
	"testing"
)

func Test_aPlusBPlusC(t *testing.T) {
	type args struct {
		as []int
		bs []int
		cs []int
		xs []int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "test1",
			args: args{
				as: []int{1, 2, 3},
				bs: []int{2, 4},
				cs: []int{1, 2, 4, 8, 16, 32},
				xs: []int{1, 5, 10, 50},
			},
			want: []bool{
				false, true, true, false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := aPlusBPlusC(tt.args.as, tt.args.bs, tt.args.cs, tt.args.xs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("aPlusBPlusC() = %v, want %v", got, tt.want)
			}
		})
	}
}
