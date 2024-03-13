package main

func main() {}

func longestPalindrome(s string) string {

	r := []rune(s)
	aheadIndex := (len(r) / 2) - 1
	var backIndex int
	var answer string
	if len(r)%2 == 0 {
		backIndex = len(r) / 2
	} else {
		backIndex = (len(r) / 2) + 1
		answer = string(r[len(r)/2])
	}

	for r[aheadIndex] == r[backIndex] {
		answer = string(r[aheadIndex]) + answer + string(r[backIndex])
		aheadIndex--
		backIndex++
	}

	return answer
}
