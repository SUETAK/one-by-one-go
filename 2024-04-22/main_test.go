package main

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{s: "Hello World"},
			want: 5,
		}, {
			name: "test",
			args: args{s: "   fly me   to   the moon  "},
			want: 4,
		}, {
			name: "test",
			args: args{s: "luffy is still joyboy"},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		}, {
			name: "test",
			args: args{strs: []string{"flower", "flow", "fright"}},
			want: "f",
		}, {
			name: "test",
			args: args{strs: []string{"flower", "frow", "flight"}},
			want: "f",
		}, {
			name: "test",
			args: args{strs: []string{"flower", "flower", "flower"}},
			want: "flower",
		}, {
			name: "test",
			args: args{strs: []string{"a"}},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
