package main

import (
	"fmt"
)

func stringInput() string {
	var s string
	fmt.Scanf("%s", &s)

	return s
}

func main() {
	s := stringInput()
	result := leftrightarrow(s)
	if result {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

func leftrightarrow(s string) bool {
	if len(s) < 3 {
		return false
	}

	if string(s[0]) != "<" {
		return false
	}
	if string(s[len(s)-1]) != ">" {
		return false
	}
	// >===< を条件から削除
	// <=> 以上であることを確定

	arrowMap := map[string]int{}

	for _, v := range s {
		if _, ok := arrowMap[string(v)]; ok {
			arrowMap[string(v)]++
			continue
		}
		arrowMap[string(v)] = 1
	}
	if arrowMap["<"] > 1 {
		return false
	}
	if arrowMap[">"] > 1 {
		return false
	}
	if arrowMap["="] <= 0 {
		return false
	}
	if len(arrowMap) > 3 {
		return false
	}
	return true
}
