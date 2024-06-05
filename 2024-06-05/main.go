package main

import "sync"

func main() {}

func isIsomorphic(s string, t string) bool {

	sCountMap := make(map[int]int)
	tCountMap := make(map[int]int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		sCountMap = convert(s)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		tCountMap = convert(t)
	}()

	wg.Wait()

	for i, _ := range sCountMap {
		if _, ok := tCountMap[i]; !ok {
			return false
		}
	}

	for i, _ := range tCountMap {
		if _, ok := sCountMap[i]; !ok {
			return false
		}
	}
	return true
}

func convert(s string) map[int]int {
	sMap := make(map[rune]int)
	countMap := make(map[int]int)
	for _, v := range s {
		if _, ok := sMap[v]; ok {
			sMap[v]++
		} else {
			sMap[v] = 1
		}
	}

	for _, i := range sMap {
		if _, ok := countMap[i]; ok {
			countMap[i]++
		} else {
			countMap[i] = 1
		}
	}

	return countMap
}
