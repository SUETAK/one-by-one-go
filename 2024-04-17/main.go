package main

import "strings"

func main() {}

func romanToInt(s string) int {
	// M, C, M, X, C, I, V
	romanCostMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	if len(s) == 1 {
		return romanCostMap[s]
	}

	split := strings.Split(s, "")
	//romanMap := make(map[string]int)

	var answer int

	before := ""
	for i := 1; i < len(split); i++ {
		before = split[i-1]
		if before == split[i] {
			answer += romanCostMap[split[i]]
		}

	}

	return answer
}
