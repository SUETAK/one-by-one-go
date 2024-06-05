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

	// 窓の右枠はnums の最大幅
	for right := 0; right < n; right++ {
		sum += nums[right]
		// sumがtarget を上回っている間だけ、leftを更新し続ける
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
