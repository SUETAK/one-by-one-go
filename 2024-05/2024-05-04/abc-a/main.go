package main

import "fmt"

func main() {
	_, x, y, z := decideInputAmountSeparatedByHalfSpace()

	result := isIncludeZ(x, y, z)

	if result {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func isIncludeZ(x, y, z int) bool {
	if x == y {
		return x == z
	}

	if x > y {
		// y < z < x
		if y <= z && z <= x {
			return true
		} else {
			return false
		}
	} else {
		// x < z < y
		if x <= z && z <= y {
			return true
		} else {
			return false
		}
	}
}

func decideInputAmountSeparatedByHalfSpace() (int, int, int, int) {
	// 1文字ずつデータ型を指定して受け取る
	var a, b, c, d int                       // int型の変数を宣言
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d) // %dでint型を代入

	return a, b, c, d
}
