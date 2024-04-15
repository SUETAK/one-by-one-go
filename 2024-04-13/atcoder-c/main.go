package main

import (
	"fmt"
	"strings"
)

func getInput() (string, string) {

	var s string
	var t string
	fmt.Scan(&s)
	fmt.Scan(&t)

	return s, t
}

func main() {
	s, t := getInput()

	if isAirportCode(s, t) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

// s の長さ3の部分列をとり、それを大文字にしたものがt
// s の長さ2 の部分列をとり、それを英大文字に変換して、末尾にXを追加したもの
// t の文字を見つけたら、tの次の文字があるかどうかを、tの文字を見つけたindex より後ろを探す
func isAirportCode(s, t string) bool {
	s = s + "x"

	t = strings.ToLower(t)
	idx := 0
	for _, r := range t {
		index := strings.Index(s[idx:], string(r))
		if index == -1 {
			return false
		}

		idx += index + 1
	}

	return true
}
