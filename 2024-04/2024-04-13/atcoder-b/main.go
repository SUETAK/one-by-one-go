package main

import "fmt"

func decideInputAmountSeparatedByHalfSpace() string {
	var s string
	fmt.Scanf("%s", &s) // %sでstring型を代入

	return s
}

func main() {
	s := decideInputAmountSeparatedByHalfSpace()
	answer := getSubstrings(s)
	if answer {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func getSubstrings(s string) bool {
	findOutRuneMap := make(map[rune]int)
	indexMap := make(map[int]int)

	for _, r := range s {
		if ok := findOutRuneMap[r]; ok > 0 {
			findOutRuneMap[r]++
		} else {
			findOutRuneMap[r] = 1
		}
	}

	for _, count := range findOutRuneMap {
		if ok := indexMap[count]; ok > 0 {
			if indexMap[count] == 2 {
				return false
			}
			indexMap[count]++
		} else {
			indexMap[count] = 1
		}
	}

	for _, count := range indexMap {
		if count == 1 {
			return false
		}
	}
	return true

}
