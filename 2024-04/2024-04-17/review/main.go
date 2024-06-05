package review

func romanToInt(s string) int {

	romanCostMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	beforeValue := romanCostMap[rune(s[len(s)-1])]
	answer := beforeValue
	for i := len(s) - 2; i >= 0; i-- {
		value := romanCostMap[rune(s[i])]
		if value >= beforeValue {
			answer += value
		} else {
			answer -= value
		}
		beforeValue = value
	}

	return answer
}
