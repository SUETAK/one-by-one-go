package main

import (
	"fmt"
)

func main() {
	// 入力
	var H, W int
	fmt.Scan(&H, &W)

	A := make([][]int, H)
	for i := 0; i < H; i++ {
		A[i] = make([]int, W)
		for j := 0; j < W; j++ {
			fmt.Scan(&A[i][j])
		}
	}

	// 前処理
	yoko := make([]int, H)
	tate := make([]int, W)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			yoko[i] += A[i][j]
			tate[j] += A[i][j]
		}
	}

	// 各マス
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Printf("%d ", yoko[i]+tate[j]-A[i][j])
		}
		fmt.Println()
	}
}
