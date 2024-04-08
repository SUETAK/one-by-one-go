package main

func main() {}

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	answer[0] = 1
	answer[len(nums)-1] = 1

	// 各要素の一つ左の要素を掛け算する
	for i := 1; i < len(nums); i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	rightProduct := 1
	for i := len(nums) - 2; i >= 0; i-- {
		rightProduct *= nums[i+1]
		answer[i] *= rightProduct
	}
	return answer
}
