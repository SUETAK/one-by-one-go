package atcoder

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mainC() {
	_, ps, bs, cs, xs := decideInputCountAndNotDecideInputAmountSeparatedBySpace()
	result := aPlusBPlusC(ps, bs, cs)
	for _, x := range xs {
		if _, ok := result[x]; ok {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func aPlusBPlusC(as, bs, cs []int) map[int]int {
	// xi の値をas, bs, cs のうちどれかの要素で表現できるかどうかを判定する
	// 100*100*100 を1回行って、作成した値をset に入れる。答えとの称号をset を使って判定する
	var valueMap = make(map[int]int)

	// 絶対に1つは選ぶ必要がある
	for _, av := range as {
		for _, bv := range bs {
			for _, cv := range cs {
				valueMap[av+bv+cv] = av + bv + cv
			}
		}

	}
	return valueMap
}

func decideInputCountAndNotDecideInputAmountSeparatedBySpace() (int, []int, []int, []int, []int) {
	var n int
	fmt.Scanf("%d", &n) // %sでstring型を代入
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")

	// 配列inputsの中身をstring→intに変換してリストに格納
	var ps []int                   // Pnを格納する配列を宣言
	for _, input := range inputs { // 配列inputsの全ての要素について実行
		p, _ := strconv.Atoi(input) // string→intに型変換
		ps = append(ps, p)          // intに型変換した値を、Pnを格納する配列に追加
	}

	var m int
	fmt.Scanf("%d", &m) // %sでstring型を代入
	sc.Scan()
	inputs = strings.Split(sc.Text(), " ")

	// 配列inputsの中身をstring→intに変換してリストに格納
	var bs []int                   // Pnを格納する配列を宣言
	for _, input := range inputs { // 配列inputsの全ての要素について実行
		p, _ := strconv.Atoi(input) // string→intに型変換
		bs = append(bs, p)          // intに型変換した値を、Pnを格納する配列に追加
	}

	var l int
	fmt.Scanf("%d", &l) // %sでstring型を代入
	sc.Scan()
	inputs = strings.Split(sc.Text(), " ")

	// 配列inputsの中身をstring→intに変換してリストに格納
	var cs []int                   // Pnを格納する配列を宣言
	for _, input := range inputs { // 配列inputsの全ての要素について実行
		p, _ := strconv.Atoi(input) // string→intに型変換
		cs = append(cs, p)          // intに型変換した値を、Pnを格納する配列に追加
	}

	var q int
	fmt.Scanf("%d", &q) // %sでstring型を代入
	sc.Scan()
	inputs = strings.Split(sc.Text(), " ")

	// 配列inputsの中身をstring→intに変換してリストに格納
	var xs []int                   // Pnを格納する配列を宣言
	for _, input := range inputs { // 配列inputsの全ての要素について実行
		p, _ := strconv.Atoi(input) // string→intに型変換
		xs = append(xs, p)          // intに型変換した値を、Pnを格納する配列に追加
	}
	return q, ps, bs, cs, xs
}
