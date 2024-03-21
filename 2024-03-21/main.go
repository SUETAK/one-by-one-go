package main

func main() {}

func removeDuplicates(nums []int) int {
	//k := 1
	//count := 1
	//previousValue := nums[0]
	//for i := 1; i < len(nums); i++ {
	//
	//}
	//return k
	return 1
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
