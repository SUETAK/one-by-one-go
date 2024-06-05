package main

import "testing"

func Test_isIncludeZ(t *testing.T) {
	type args struct {
		x int
		y int
		z int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test",
			args: args{
				x: 0,
				y: 0,
				z: 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIncludeZ(tt.args.x, tt.args.y, tt.args.z); got != tt.want {
				t.Errorf("isIncludeZ() = %v, want %v", got, tt.want)
			}
		})
	}
}
