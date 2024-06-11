package main

import "strings"

func main() {}

func wordPattern(pattern string, s string) bool {
	sArray := strings.Split(s, " ")
	pMap := make(map[rune]string)
	sMap := make(map[string]rune)

	if len(sArray) != len(pattern) {
		return false
	}
	for i, pr := range pattern {
		sVal, isContainS := pMap[pr]
		prVal, isContainPR := sMap[sArray[i]]

		if !isContainS && !isContainPR {
			pMap[pr] = sArray[i]
			sMap[sArray[i]] = pr
		} else if isContainS && sVal != sArray[i] {
			return false
		} else if isContainPR && prVal != pr {
			return false
		}
	}
	return true
}
