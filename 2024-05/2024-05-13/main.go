package main

import "sort"

func main() {}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := nums[i]
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := target + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{target, nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
				continue
			}
			if sum < 0 {
				left++
				continue
			}

			right--
			continue

		}
	}
	return result
}
