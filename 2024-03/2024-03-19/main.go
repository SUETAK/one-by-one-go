package main

func main() {}

func removeElement(nums []int, val int) int {
	var answer int
	for _, v := range nums {
		if v != val {
			nums[answer] = v
			answer++
		}
	}

	return answer
}
