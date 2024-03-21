package main

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want2 []int
	}{
		{
			name:  "test1",
			args:  args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3}},
			want:  7,
			want2: []int{0, 0, 1, 1, 2, 2, 3, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := tt.args.nums
			if got := removeDuplicates(nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
			for i, v := range nums[:5] {
				if v != tt.want2[i] {
					t.Errorf("removeDuplicates() = %v, want %v", v, tt.want2[i])
				}
			}
		})
	}
}
