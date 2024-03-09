package atcoder

import "testing"

func Test_spoiler(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{s: "abc|def|ghi"},
			want: "abcghi",
		},
		{
			name: "test2",
			args: args{s: "|spoiler|"},
			want: "",
		},
		{
			name: "test3",
			args: args{s: "||spoiler"},
			want: "spoiler",
		},
		{
			name: "test4",
			args: args{s: "spoiler||"},
			want: "spoiler",
		},
		{
			name: "test5",
			args: args{s: "spoiler|test|"},
			want: "spoiler",
		},
		{
			name: "test6",
			args: args{s: "|test|spoiler"},
			want: "spoiler",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spoiler(tt.args.s); got != tt.want {
				t.Errorf("spoiler() = %v, want %v", got, tt.want)
			}
		})
	}
}
