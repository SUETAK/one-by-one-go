package main

import (
	"fmt"
	"strings"
)

func decideInputAmountSeparatedByHalfSpace() (string, string) {

	var s string
	var t string
	fmt.Scan(&s)
	fmt.Scan(&t)

	return s, t
}

func main() {
	s, t := decideInputAmountSeparatedByHalfSpace()

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
	ta := strings.Split(strings.ToLower(t), "")
	idx := 0
	for _, r := range s {
		if ta[idx] == string(r) {
			idx += 1
		}
		if idx == 3 {
			return true
		}
	}
	if idx == 2 && ta[idx] == "x" {
		return true
	}
	return false
}
