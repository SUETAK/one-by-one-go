package main

import "strings"

func main() {}

var romanStringCostMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func intToRoman(num int) string {

	answer := ""
	for num > 0 {
		if num/1000 != 0 {
			num, answer = decreaseNumber(num, "M", "M", "M", answer)
		}
		if num != 0 && num/500 != 0 {
			num, answer = decreaseNumber(num, "C", "D", "M", answer)
		}
		if num != 0 && num/100 != 0 {
			num, answer = decreaseNumber(num, "L", "C", "D", answer)
		}
		if num != 0 && num/50 != 0 {
			num, answer = decreaseNumber(num, "X", "L", "C", answer)
		}
		if num != 0 && num/10 != 0 {
			num, answer = decreaseNumber(num, "V", "X", "L", answer)
		}
		if num != 0 && num/5 != 0 {
			num, answer = decreaseNumber(num, "I", "V", "X", answer)
		}
		if num != 0 && num/1 != 0 {
			num, answer = decreaseNumber(num, "I", "I", "V", answer)
		}
	}

	return answer
}

func decreaseNumber(num int, beforeDigit, digit, nextDigit, answer string) (int, string) {
	cost := romanStringCostMap[digit]
	if num/romanStringCostMap[beforeDigit] == 9 {
		answer += beforeDigit + nextDigit
		num -= romanStringCostMap[beforeDigit] * 9
		return num, answer
	} else if num/romanStringCostMap[digit] == 4 {
		answer += digit + nextDigit
	} else {
		answer += strings.Repeat(digit, num/cost)
	}
	if cost == 1 {
		num -= num / cost
	} else {
		num -= cost * (num / cost)
	}
	return num, answer
}
