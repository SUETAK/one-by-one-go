package main

func main() {}

func majorityElement(nums []int) int {
	elementMap := map[int]int{}
	majorCount := len(nums)/2 + 1
	var result int
	for _, v := range nums {
		elementMap[v]++
		if elementMap[v] == majorCount {
			result = v
			break
		}
	}

	return result
}
