package main

import "fmt"

func twoInputAmountSeparatedBySpace() (int, int) {
	// 1文字ずつデータ型を指定して受け取る
	var a, b int                  // int型の変数を宣言
	fmt.Scanf("%d %d %d", &a, &b) // %dでint型を代入

	return a, b
}

func main() {
	w, b := twoInputAmountSeparatedBySpace()

}

func findCombination(w, b int) bool {
	base := "wbwbwwbwbwbw"
	conbMap := map[string]int{"w": 0, "b": 0}
	for conbMap["w"] > w || conbMap["b"] > b {
		for _, v := range base {
			conbMap[string(v)]++
		}
	}

}
