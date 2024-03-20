package main

func main() {}

func removeDuplicates(nums []int) int {
	k := 1
	previousValue := nums[0]
	for _, v := range nums {
		if v != previousValue {
			nums[k] = v
			k++
			previousValue = v
		}
	}
	return k
}
