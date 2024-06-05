package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {}

func isPalindrome(s string) bool {
	filtered := strings.Join(strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	}), "")
	lower := strings.ToLower(filtered)
	fmt.Println(lower)

	k := len(lower) - 1
	for i := 0; i < k; i++ {
		head := lower[i]
		tail := lower[k]
		if head == tail {
			k--
			continue
		} else {
			return false
		}
	}
	return true
}
