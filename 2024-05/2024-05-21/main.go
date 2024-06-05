package main

import "strings"

func main() {}

func lengthOfLongestSubstring(s string) int {

	maxLength := 0
	left := 0

	for right := 0; right < len(s); right++ {
		currentLength := right - left + 1
		if currentLength > maxLength {
			if i := strings.IndexByte(s[left:right], s[right]); i >= 0 {

				left = i + left + 1
			} else {
				maxLength = currentLength
			}
		}
	}
	return maxLength
}
