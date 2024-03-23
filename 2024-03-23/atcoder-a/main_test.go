package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func Test_calc(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{nums: []int{3, 4, 6}},
			want: []string{"12", "24"},
		},
		{
			name: "test1",
			args: args{nums: []int{22, 75, 26, 45, 72}},
			want: []string{"1650", "1950", "1170", "3240"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calc(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calc() = %v, want %v", got, tt.want)
			}
			fmt.Println(strings.Join(calc(tt.args.nums), " "))
		})
	}
}
