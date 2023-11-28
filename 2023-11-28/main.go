package main

import "sort"

func main() {}

// https://www.hackerrank.com/challenges/migratory-birds/problem
func migratoryBirds(arr []int32) int32 {

	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	bardsMap := make(map[int32]int32)
	for _, v := range arr {
		if _, ok := bardsMap[v]; ok {
			bardsMap[v]++
		} else {
			bardsMap[v] = 1
		}
	}

	var maxBardType int32
	var maxCount int32
	for key, v := range bardsMap {
		if maxCount < v {
			maxCount = v
			maxBardType = key
		}
	}
	return maxBardType
}
