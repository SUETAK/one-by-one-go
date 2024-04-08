package main

func main() {}

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	answer[0] = 1
	answer[len(nums)-1] = 1

	// 各要素の一つ左の要素を掛け算する
	// 一番左は左がないのでi=1 から始める
	for i := 1; i < len(nums); i++ {
		// answer のi-1(1つ左)は左側の累積になっている
		answer[i] = answer[i-1] * nums[i-1]
	}

	rightProduct := 1
	// len(nums)-1 が配列の最後のindex
	// 一番右より右はないので最後のindex から-1 する
	lastIndex := len(nums) - 1
	for i := lastIndex - 1; i >= 0; i-- {
		rightProduct *= nums[i+1] // 累積
		answer[i] *= rightProduct // 左側の累積*右側の累積
	}
	return answer
}
