package main

func main() {}

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 0
	}
	n := len(nums)

	count := 0
	maxReach := 0
	for i := 0; i < len(nums); i++ {
		maxReach += nums[i]
		if maxReach >= n-1 {
			count++
			break
		}
		count++
	}
	return count
}
