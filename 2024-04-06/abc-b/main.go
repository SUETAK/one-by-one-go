package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func NotDecideInputAmountMultiColumns() [][]int {
	// 1文字ずつデータ型を指定して受け取る
	var n int           // int型の変数を宣言
	fmt.Scanf("%d", &n) // %dでint型を代入

	sc := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ

	var us [][]int // uの値を格納する配列を宣言

	for i := 0; i < n; i++ { // 行数分繰り返す
		sc.Scan()                               // １行分の入力を取得する
		inputs := strings.Split(sc.Text(), " ") // 半角スペース区切りでstring型として配列inputsに格納
		x, _ := strconv.Atoi(inputs[0])         // string→intに型変換
		y, _ := strconv.Atoi(inputs[1])         // string→intに型変換
		us = append(us, []int{x, y})            // uの配列に追加
	}
	return us
}

func main() {
	dots := NotDecideInputAmountMultiColumns()
	result := mostFarDot(dots)
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func mostFarDot(dots [][]int) []int {
	result := make([]int, 0, len(dots))

	for i := 0; i < len(dots); i++ {
		curDot := dots[i]
		farDotIndex := 0
		mostFarDistance := float64(0)
		for j := 0; j < len(dots); j++ {
			if i == j {
				continue
			}
			d := distance(float64(curDot[0]), float64(curDot[1]), float64(dots[j][0]), float64(dots[j][1]))

			if mostFarDistance < d {
				farDotIndex = j + 1
				mostFarDistance = d
			}
		}
		result = append(result, farDotIndex)
	}

	return result
}
