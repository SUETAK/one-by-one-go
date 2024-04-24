package main

import "strings"

func main() {}

func reverseWords(s string) string {
	split := strings.Split(s, " ")
	result := make([]string, 0)
	for i := len(split) - 1; i >= 0; i-- {
		if split[i] != "" {
			result = append(result, split[i])
		}
	}
	return strings.Join(result, " ")
}
