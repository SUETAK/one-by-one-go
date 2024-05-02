package main

import "strings"

func main() {}

func canConstruct(ransomNote string, magazine string) bool {
	rMap := make(map[string]int)

	for _, v := range ransomNote {
		if _, ok := rMap[string(v)]; !ok {
			rMap[string(v)] = 1
		} else {
			rMap[string(v)]++
		}
	}

	for s, i := range rMap {
		count := strings.Count(magazine, s)
		if count < i {
			return false
		}
	}
	return true
}
