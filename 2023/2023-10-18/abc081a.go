package main

func countPutOnBallCells(a, b, c int) int {
	inputSlice := []int{a, b, c}
	var count int
	for _, input := range inputSlice {
		if input == 1 {
			count++
		}
	}
	return count
}
