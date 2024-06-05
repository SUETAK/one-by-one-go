package main

import "sort"

func main() {}

func hIndex(citations []int) int {

	if len(citations) == 1 {
		if citations[0] == 0 {
			return 0
		} else {
			return 1
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	for i := 0; i < len(citations); i++ {
		if i >= citations[i] {
			return i
		}
	}
	return len(citations)
}
