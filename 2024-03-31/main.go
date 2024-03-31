package main

func main() {}

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	maxReach := nums[0]
	for i := 1; i < len(nums); i++ {
		if maxReach < i {
			return false
		}
		if maxReach > len(nums) {
			return true
		}
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
	}
	return true
}
