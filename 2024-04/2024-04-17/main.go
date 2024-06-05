package main

import "strings"

func main() {}

func romanToInt(s string) int {
	romanCostMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	split := strings.Split(s, "")

	var answer, currentValue, lastValue int

	for i := len(s) - 1; i >= 0; i-- {
		currentValue = romanCostMap[split[i]]
		if currentValue < lastValue {
			answer -= currentValue
		} else {
			answer += currentValue
		}
		// 前回の値の更新
		lastValue = currentValue
	}

	return answer
}
