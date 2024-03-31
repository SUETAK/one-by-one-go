package main

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := tt.args.nums
			rotate(nums, tt.args.k)
			for i, v := range nums {
				if v != tt.want[i] {
					t.Errorf("removeDuplicates() = %v, want %v", v, tt.want[i])
				}
			}
		})
	}
}
