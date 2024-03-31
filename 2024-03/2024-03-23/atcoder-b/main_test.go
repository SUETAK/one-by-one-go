package main

import "testing"

func Test_findCombination(t *testing.T) {
	type args struct {
		w int
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1",
			args: args{
				w: 3,
				b: 2,
			},
			want: true,
		},
		{name: "test2",
			args: args{
				w: 3,
				b: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCombination(tt.args.w, tt.args.b); got != tt.want {
				t.Errorf("findCombination() = %v, want %v", got, tt.want)
			}
		})
	}
}
