package main

func main() {}

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	if nums[0] == 0 {
		return false
	}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] != 0 {
			continue
		}

		// 0があったら前をたどる
		count := 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] <= count {
				if j == 0 {
					return false
				}
				count++
				continue
			} else {
				break
			}
		}

	}
	return true
}
