package main

func main() {}

func minSubArrayLen(target int, nums []int) int {

	left := 0
	right := 0
	result := len(nums)
	sum := nums[right]
	for left != len(nums) {
		for right != len(nums)-1 && sum < target {
			right++
			sum += nums[right]

			if sum == target {
				length := right + 1 - left
				if result > length {
					result = length
				}
				left = right
				sum = 0
				continue
			}

		}

		sum -= nums[left]
		left++
		if sum == target {
			length := right + 1 - left
			if result > length {
				result = length
			}
			left = right
			sum = 0
			continue
		}
		// sum > target の状況
		for sum > target || len(nums)-1 == right {
			left++
			sum -= nums[left]

			if sum == target {
				length := right + 1 - left
				if result > length {
					result = length
				}
				left = right
				sum = 0
				continue
			}
		}
	}

	return result
}
