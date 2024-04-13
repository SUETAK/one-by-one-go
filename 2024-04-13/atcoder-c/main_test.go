package main

import "testing"

func Test_isAirportCode(t *testing.T) {
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
				s: "narita",
				t: "NRT",
			},
			want: true,
		}, {
			name: "test",
			args: args{
				s: "losangeles",
				t: "LAX",
			},
			want: true,
		}, {
			name: "test",
			args: args{
				s: "snuke",
				t: "RNG",
			},
			want: false,
		}, {
			name: "test",
			args: args{
				s: "snuke",
				t: "RNG",
			},
			want: false,
		}, {
			name: "test",
			args: args{
				s: "hasvikairport",
				t: "HAA",
			},
			want: true,
		}, {
			name: "test",
			args: args{
				s: "abcde",
				t: "BCX",
			},
			want: true,
		}, {
			name: "test",
			args: args{
				s: "abcde",
				t: "XYZ",
			},
			want: false,
		}, {
			name: "test",
			args: args{
				s: "cbacb",
				t: "ABC",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAirportCode(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAirportCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
