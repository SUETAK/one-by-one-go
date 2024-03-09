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
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "test1",
			args: args{
				as: []int{1, 2},
				bs: []int{2, 3},
				cs: []int{3, 4},
			},
			want: map[int]int{
				6: 6, 7: 7, 8: 8, 9: 9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := aPlusBPlusC(tt.args.as, tt.args.bs, tt.args.cs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("aPlusBPlusC() = %v, want %v", got, tt.want)
			}
		})
	}
}
