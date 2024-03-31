package main

import "fmt"

func stringInput() string {
	var s string
	fmt.Scanf("%s", &s)

	return s
}
func main() {
	s := stringInput()
	fmt.Println(oneTimeSwap(s))
}

func oneTimeSwap(s string) int {

	if len(s) < 2 {
		return 1
	}

	stringMap := map[rune]int{}
	for _, v := range s {
		stringMap[v]++
	}

	sameCombinationCount := 0
	// 重ならないやつ
	for _, v := range stringMap {
		sameCombinationCount += getCombinationCount(v)
	}
	answer := getCombinationCount(len(s)) - sameCombinationCount

	if sameCombinationCount > 0 {
		answer++
	}

	return answer
}

func getCombinationCount(n int) int {
	return n * (n - 1) / 2
}
