package main

import "testing"

func Test_higherIndex(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{input: []int{3, 2, 5, 2}},
			want: 3,
		}, {
			name: "test",
			args: args{input: []int{4, 3, 2}},
			want: -1,
		}, {
			name: "test",
			args: args{input: []int{10, 5, 10, 2, 10, 13, 15}},
			want: 6,
		}, {
			name: "test",
			args: args{input: []int{10}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := higherIndex(tt.args.input); got != tt.want {
				t.Errorf("higherIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
