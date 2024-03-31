package main

import "testing"

func Test_integerDivisionReturns(t *testing.T) {
	type args struct {
		s int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				s: 27,
			},
			want: 3,
		},
		{
			name: "test1",
			args: args{
				s: -13,
			},
			want: -1,
		},
		{
			name: "test1",
			args: args{
				s: 40,
			},
			want: 4,
		},
		{
			name: "test1",
			args: args{
				s: -20,
			},
			want: -2,
		},
		{
			name: "test1",
			args: args{
				s: 123456789123456789,
			},
			want: 12345678912345679,
		},
		{
			name: "test1",
			args: args{
				s: 1,
			},
			want: 1,
		},
		{
			name: "test1",
			args: args{
				s: -1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerDivisionReturns(tt.args.s); got != tt.want {
				t.Errorf("integerDivisionReturns() = %v, want %v", got, tt.want)
			}
		})
	}
}
