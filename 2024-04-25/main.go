package main

import "strings"

func main() {}

func strStr(haystack string, needle string) int {
	result := strings.Index(haystack, needle)
	return result
}
