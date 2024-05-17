package main

import "math"

func main() {}

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	minLength := len(nums)
	sum := 0
	left := 0

	for right := 0; right < n; right++ {
		sum += nums[right]
		// Once sum is greater than or equal to target, try to shrink the window size
		for sum >= target {
			if currentLength := right - left + 1; currentLength < minLength {
				minLength = currentLength
			}
			sum -= nums[left]
			left++
		}
	}

	if minLength == math.MaxInt32 {
		return 0
	}
	return minLength
}
