package main

import (
	"sort"
	"strings"
)

func main() {}

func lengthOfLastWord(s string) int {
	split := strings.Split(s, " ")
	filteredSplit := make([]string, 0)
	for _, v := range split {
		if len(v) != 0 {
			filteredSplit = append(filteredSplit, v)
		}
	}

	return len(filteredSplit[len(filteredSplit)-1])
}

func longestCommonPrefix(strs []string) string {
	result := ""
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	targetWord := strs[0]
	for i := 0; i != len(targetWord); i++ {
		targetString := targetWord[i]
		for _, v := range strs {
			if targetString != v[i] {
				return result
			}
		}
		result += string(targetString)
	}
	return result
}
