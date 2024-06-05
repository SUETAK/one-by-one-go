package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func NotDecideInputAmountMultiColumns() int {
	// 1文字ずつデータ型を指定して受け取る
	var n int           // int型の変数を宣言
	fmt.Scanf("%d", &n) // %dでint型を代入

	sc := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ

	colorMinDeliciousValueMap := make(map[int]int, n) //色 : 最小の値
	for i := 0; i < n; i++ {
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")
		a, _ := strconv.Atoi(inputs[0])
		c, _ := strconv.Atoi(inputs[1])

		if _, ok := colorMinDeliciousValueMap[c]; !ok {
			colorMinDeliciousValueMap[c] = a
		}
		if a < colorMinDeliciousValueMap[c] {
			colorMinDeliciousValueMap[c] = a
		}

	}

	maxValue := 0
	for _, v := range colorMinDeliciousValueMap {
		if v > maxValue {
			maxValue = v
		}

	}
	return maxValue
}

func main() {
	fmt.Println(NotDecideInputAmountMultiColumns())
}
