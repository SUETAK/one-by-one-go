package main

func main() {}

func rotate(nums []int, k int) {
	if k%len(nums) == 0 {
		return
	}
	separate := k % len(nums)
	right := nums[separate+1:]
	left := nums[:separate+1]
	result := append(right, left...)
	for i, v := range result {
		nums[i] = v
	}
}
