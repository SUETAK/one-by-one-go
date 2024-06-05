package main

import "fmt"

func main() {
	var n int           // int型の変数を宣言
	fmt.Scanf("%d", &n) // %dでint型を代入

	fmt.Println(getKickResult(n))
}

func getKickResult(n int) string {
	result := ""
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			result = result + "x"
		} else {
			result = result + "o"
		}
	}
	return result
}
