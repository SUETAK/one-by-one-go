package main

import (
	"fmt"
)

func notDecideInputAmountSeparatedBySpace() ([]int, int) {
	var n int    // int型の変数を宣言
	fmt.Scan(&n) // %1dでint型を1つずつ代入代入

	scores := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Scan(&scores[i])
	}
	return scores, n
}

func main() {
	ps, n := notDecideInputAmountSeparatedBySpace()
	answer := getWinnersPoint(n, ps)
	fmt.Println(answer)
}

// ps でn-1 人目までの点数が規定されていて、n人目の点数を返す
// 合計は常に0 になるので、全体を足して、合計が0 になる数を返す
func getWinnersPoint(n int, ps []int) int {
	sum := 0
	for _, p := range ps {
		sum += p
	}
	return -sum
}
