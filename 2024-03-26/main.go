package main

func main() {}

func rotate(nums []int, k int) {
	separate := k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, separate-1)
	reverse(nums, k, len(nums)-1)
}

// reverse 与えられた配列を、start~end で逆順にする
func reverse(nums []int, start int, end int) {
	for start <= end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func rotateBest(nums []int, k int) {
	num := len(nums)
	k = k % num
	temp := make([]int, 0, num)
	// num-k を分割する箇所として考えられればOKだった
	left := append(temp, nums[num-k:]...)
	temp = append(left, nums[:num-k]...)

	copy(nums, temp)
}
