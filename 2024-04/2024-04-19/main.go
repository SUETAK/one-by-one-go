package main

import "strings"

func main() {}

var romanStringCostMap = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func intToRoman(num int) string {
	answer := ""
	numArray := make(map[int]int)
	r := []int{1000, 100, 10, 1}
	for _, v := range r {

		numArray[v] = num / v
		num -= v * (num / v)
	}

	for _, rv := range r {
		if numArray[rv] == 0 {
			continue
		}
		if rv == 1000 {
			answer += strings.Repeat(romanStringCostMap[rv], numArray[rv])
			continue
		}

		if numArray[rv] == 4 || numArray[rv] == 5 || numArray[rv] == 9 {
			answer += romanStringCostMap[numArray[rv]*rv]
			continue
		}

		if numArray[rv] <= 3 {
			answer += strings.Repeat(romanStringCostMap[rv], numArray[rv])
			continue
		}
		if numArray[rv] <= 8 {
			answer += romanStringCostMap[rv*5] + strings.Repeat(romanStringCostMap[rv], numArray[rv]-5)
			continue
		}
	}
	return answer
}
