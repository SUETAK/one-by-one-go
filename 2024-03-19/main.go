package main

func main() {}

func removeElement(nums []int, val int) int {
	var removedNums []int
	for _, v := range nums {
		if v != val {
			removedNums = append(removedNums, v)
		}
	}
	nums = removedNums
	return len(removedNums)
}
