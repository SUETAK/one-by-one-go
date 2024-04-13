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
func isAirportCode(s, t string) bool {
	lowerT := strings.ToLower(t)
	tMap := map[string]bool{
		lowerT[0:1]: false,
		lowerT[1:2]: false,
		lowerT[2:3]: lowerT[2:3] == "x",
	}
	index := 0
	for _, r := range lowerT {
		findIndex := strings.Index(s[index:], string(r))
		if findIndex != -1 {
			if index < findIndex {
				index = findIndex
			} else if index == findIndex {
				index++
			}
			tMap[string(r)] = true
		}
	}

	for _, v := range tMap {
		if !v {
			return false
		}
	}
	return true
}
