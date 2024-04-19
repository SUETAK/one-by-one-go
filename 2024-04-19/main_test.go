package main

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{num: 3},
			want: "III",
		}, {
			name: "test",
			args: args{num: 4},
			want: "IV",
		}, {
			name: "test",
			args: args{num: 5},
			want: "V",
		}, {
			name: "test",
			args: args{num: 58},
			want: "LVIII",
		}, {
			name: "test",
			args: args{num: 50},
			want: "L",
		}, {
			name: "test",
			args: args{num: 20},
			want: "XX",
		}, {
			name: "test",
			args: args{num: 100},
			want: "C",
		}, {
			name: "test",
			args: args{num: 500},
			want: "D",
		}, {
			name: "test",
			args: args{num: 1000},
			want: "M",
		}, {
			name: "test",
			args: args{num: 45},
			want: "XLV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
