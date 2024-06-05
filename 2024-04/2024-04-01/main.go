package main

func main() {}

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	maxRange := 0
	for i := 0; i < len(nums); i++ {
		if maxRange <= i+nums[i] {
			maxRange = i + nums[i]
		}
		if len(nums)-1 <= maxRange {
			return true
		}
		if maxRange <= i {
			return false
		}
	}
	return false
}
