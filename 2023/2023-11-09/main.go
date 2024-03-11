package main

import "fmt"

// https://atcoder.jp/contests/typical90/tasks/typical90_a
// https://github.com/E869120/kyopro_educational_90/blob/main/editorial/001.jpg
func main() {
	// 入力 取り込み
	//var N, L, K int
	//fmt.Scan(&N, &L, &K)
	_, L, K := 7, 45, 2

	//A := make([]int, N)
	//for i := 0; i < N; i++ {
	//	fmt.Scan(&A[i])
	//}

	A := []int{7, 11, 16, 20, 28, 34, 38}
	// 二部探索
	left, right := -1, 46

	for right-left > 1 {
		mid := (left + right) / 2
		if check(L, mid, K, A) {
			left = mid
		} else {
			right = mid
		}
	}
	fmt.Println(left)

}

func check(L, mid, K int, A []int) bool {
	// 何個切れたか
	num := 0 // 何個切れたか
	pre := 0 // 前回の切れ目

	for _, v := range A {
		// mid を超えたら切断
		if v-pre >= mid {
			num++
			pre = v
		}
	}
	// 最後のピースが mid 以上なら加算
	if L-pre >= mid {
		num++
	}

	return num >= K+1
}
