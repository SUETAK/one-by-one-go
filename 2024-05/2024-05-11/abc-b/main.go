package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	_, k, a := twoInputAmountSeparatedBySpace()

	fmt.Println(countStart(k, a))
}

func countStart(k int, group []int) int {
	if len(group) == 1 {
		return 1
	}
	count := 0

	sum := 0
	for i, v := range group {
		sum += v
		if sum == k {
			count++
			sum = 0
			continue
		}

		if sum > k {
			count++
			sum = v
			if i == len(group)-1 {
				count++
			}
			continue
		}

		if sum < k {
			if i == len(group)-1 {
				count++
			}
			continue
		}
	}
	return count
}

func twoInputAmountSeparatedBySpace() (int, int, []int) {
	// 1文字ずつデータ型を指定して受け取る
	var a, b int                  // int型の変数を宣言
	fmt.Scanf("%d %d %d", &a, &b) // %dでint型を代入

	sc := bufio.NewScanner(os.Stdin)        // 標準入力を受け付けるスキャナ
	sc.Scan()                               // １行分の入力を取得する
	inputs := strings.Split(sc.Text(), " ") // 半角スペース区切りでstring型として配列inputsに格納

	// 配列inputsの中身をstring→intに変換してリストに格納
	var ps []int                   // Pnを格納する配列を宣言
	for _, input := range inputs { // 配列inputsの全ての要素について実行
		p, _ := strconv.Atoi(input) // string→intに型変換
		ps = append(ps, p)          // intに型変換した値を、Pnを格納する配列に追加
	}

	return a, b, ps
}
