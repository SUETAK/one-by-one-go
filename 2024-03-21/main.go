package main

func main() {}

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	slowPointer := 2
	for fastPointer := 2; fastPointer < len(nums); fastPointer++ {
		// fastPointer の値とslowPointer の2つ前の値が異なる場合、
		if nums[fastPointer] != nums[slowPointer-2] {
			nums[slowPointer] = nums[fastPointer]
			slowPointer++
		}
	}
	return slowPointer
}

// 回答
func removeDuplicatesAnswer(nums []int) int {
	myMap := make(map[int]int)
	k := 0

	for i := 0; i < len(nums); i++ {
		value := nums[i]
		// value: count として、value のカウントが1なら
		if myMap[value] < 2 {
			// value のカウントをインクリメント
			myMap[value] = myMap[value] + 1
			// num[0] = nums[i]
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
