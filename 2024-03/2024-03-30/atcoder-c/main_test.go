package main

import (
	"testing"
)

func Test_isAllHoliday(t *testing.T) {
	type args struct {
		a  int
		b  int
		ps []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test",
			args: args{
				a:  5,
				b:  10,
				ps: []int{10, 15},
			},
			want: false,
		},
		{
			name: "test",
			args: args{
				a:  347,
				b:  347,
				ps: []int{347, 700, 705, 710},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isAllHoliday(tt.args.a, tt.args.b, tt.args.ps)
		})
	}
}
