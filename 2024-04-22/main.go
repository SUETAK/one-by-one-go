package main

import "strings"

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
