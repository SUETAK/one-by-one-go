package atcoder

import (
	"fmt"
)

func main() {
	s := input()
	fmt.Println(spoiler(s))
}

func spoiler(s string) string {
	// | で挟まれた文字列を削除する
	// 例:abc|def|ghi -> abcghi
	// ["a", "|", "b", "|", "c
	var pipeIncenses []int
	for i, v := range s {
		if v == '|' {
			pipeIncenses = append(pipeIncenses, i)
		}
	}

	return s[:pipeIncenses[0]] + s[pipeIncenses[1]+1:]
}

func input() string {
	var s string
	fmt.Scanf("%s", &s)
	return s
}
