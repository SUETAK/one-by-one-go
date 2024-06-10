package main

import "testing"

func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test",
			args: args{
				s: "egg",
				t: "goo",
			},
			want: true,
		}, {
			name: "test",
			args: args{
				s: "foo",
				t: "bar",
			},
			want: false,
		}, {
			name: "test",
			args: args{
				s: "paper",
				t: "title",
			},
			want: true,
		}, {
			name: "test",
			args: args{
				s: "bbbaaaba",
				t: "aaabbbba",
			},
			want: false,
		}, {
			name: "test",
			args: args{
				s: "\u0000\u0000",
				t: "\u0000c",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
