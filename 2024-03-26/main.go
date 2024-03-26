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
