package main

func main() {}

func twoSum(nums []int, t int) []int {
	head := 0
	tail := len(nums) - 1
	sum := nums[head] + nums[tail]
	for sum != t {
		if sum < t {
			head++
		} else {
			tail--
		}
		sum = nums[head] + nums[tail]
	}
	return []int{head + 1, tail + 1}
}
