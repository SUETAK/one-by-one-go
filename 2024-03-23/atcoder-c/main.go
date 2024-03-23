package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetGrid() (int, int, []int) {
	sc := bufio.NewScanner(os.Stdin)
	// 入力
	var H, W int
	fmt.Scan(&H, &W)

	sc.Scan()                               // １行分の入力を取得する
	inputs := strings.Split(sc.Text(), " ") // 半角スペース区切りでstring型として配列inputsに格納

	// 配列inputsの中身をstring→intに変換してリストに格納
	var ps []int                   // Pnを格納する配列を宣言
	for _, input := range inputs { // 配列inputsの全ての要素について実行
		p, _ := strconv.Atoi(input) // string→intに型変換
		ps = append(ps, p)          // intに型変換した値を、Pnを格納する配列に追加
	}
	return H, W, ps
}

func main() {
	_, k, ary := GetGrid()
	result := sigma(k, ary)
	fmt.Println(result)
}

func sigma(k int, nums []int) int {
	countMap := map[int]bool{}
	result := (k * (k + 1)) / 2 // 総和の式

	for _, v := range nums {
		countMap[v] = true
	}

	for i := range countMap {
		if k < i {
			continue
		}
		result -= i
	}
	return result
}
