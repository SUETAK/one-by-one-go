package main

import "testing"

func Test_countStart(t *testing.T) {
	type args struct {
		k     int
		group []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				k:     5,
				group: []int{1, 2, 3, 4, 5},
			},
			want: 4,
		}, {
			name: "test",
			args: args{
				k:     6,
				group: []int{2, 5, 1, 4, 1, 2, 3},
			},
			want: 4,
		}, {
			name: "test",
			args: args{
				k:     100,
				group: []int{73, 8, 55, 26, 97, 48, 37, 47, 35, 55, 5, 17, 62, 2, 60},
			},
			want: 8,
		}, {
			name: "test",
			args: args{
				k:     1,
				group: []int{1},
			},
			want: 1,
		}, {
			name: "test",
			args: args{
				k:     3,
				group: []int{3, 1, 1, 1, 2, 2},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countStart(tt.args.k, tt.args.group); got != tt.want {
				t.Errorf("countStart() = %v, want %v", got, tt.want)
			}
		})
	}
}
