package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 入力パターン A B
// 入力パターン S
func twoInputAmountSeparatedBySpace() (int, int) {
	// 1文字ずつデータ型を指定して受け取る
	var a, b int                  // int型の変数を宣言
	fmt.Scanf("%d %d %d", &a, &b) // %dでint型を代入

	return a, b
}

// 入力パターン A B C
// 入力パターン S
func decideInputAmountSeparatedByHalfSpace() (int, int, int, string) {
	// 1文字ずつデータ型を指定して受け取る
	var a, b, c int                   // int型の変数を宣言
	var s string                      // string型の変数を宣言
	fmt.Scanf("%d %d %d", &a, &b, &c) // %dでint型を代入
	fmt.Scanf("%s", &s)               // %sでstring型を代入

	return a, b, c, s
}

// 入力パターン DEF
func decideInputAmountNotSeparated() (int, int, int) {
	// 1文字ずつデータ型を指定して受け取る
	var d, e, f int                    // int型の変数を宣言
	fmt.Scanf("%1d%1d%1d", &d, &e, &f) // %1dでint型を1つずつ代入

	return d, e, f
}

// 入力パターン
// P1 P2 P3 ... Pn
func notDecideInputAmountSeparatedBySpace() []int {
	sc := bufio.NewScanner(os.Stdin)        // 標準入力を受け付けるスキャナ
	sc.Scan()                               // １行分の入力を取得する
	inputs := strings.Split(sc.Text(), " ") // 半角スペース区切りでstring型として配列inputsに格納

	// 配列inputsの中身をstring→intに変換してリストに格納
	var ps []int                   // Pnを格納する配列を宣言
	for _, input := range inputs { // 配列inputsの全ての要素について実行
		p, _ := strconv.Atoi(input) // string→intに型変換
		ps = append(ps, p)          // intに型変換した値を、Pnを格納する配列に追加
	}
	return ps
}

// 入力パターン
// N
// P1 P2 P3 ... Pn
func decideInputCountAndNotDecideInputAmountSeparatedBySpace() (int, []int, []int, []int, []int) {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	var m int
	fmt.Fscan(in, &m)

	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i])
	}

	var l int
	fmt.Fscan(in, &l)

	c := make([]int, l)
	for i := 0; i < l; i++ {
		fmt.Fscan(in, &c[i])
	}

	var q int
	fmt.Fscan(in, &q)
	x := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &x[i])
	}
	return q, a, b, c, x
}

// NotDecideInputAmountMultiColumns 入力パターン
// N
// U1 V1
// U2 V1
// ...
// UN VN
func NotDecideInputAmountMultiColumns() ([]int, []int) {
	// 1文字ずつデータ型を指定して受け取る
	var n int           // int型の変数を宣言
	fmt.Scanf("%d", &n) // %dでint型を代入

	sc := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ

	var us []int // uの値を格納する配列を宣言
	var vs []int // vの値を格納する配列を宣言

	for i := 0; i < n; i++ { // 行数分繰り返す
		sc.Scan()                               // １行分の入力を取得する
		inputs := strings.Split(sc.Text(), " ") // 半角スペース区切りでstring型として配列inputsに格納
		u, _ := strconv.Atoi(inputs[0])         // string→intに型変換
		v, _ := strconv.Atoi(inputs[1])         // string→intに型変換
		us = append(us, u)                      // uの配列に追加
		vs = append(vs, v)                      // vの配列に追加
	}
	return us, vs
}

// NotDecideInputAmountMultiColumnsBySliceMap  入力パターン
// N
// U1 V1
// U2 V1
// ...
// UN VN
// {N: [UN, VN]}
func NotDecideInputAmountMultiColumnsBySliceMap() (int, map[int][]int) {
	// 1文字ずつデータ型を指定して受け取る
	var n int           // int型の変数を宣言
	fmt.Scanf("%d", &n) // %dでint型を代入

	sc := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ

	var inputSliceMap map[int][]int
	for i := 0; i < n; i++ { // 行数分繰り返す
		var intInputs []int
		sc.Scan()                                 // １行分の入力を取得する
		inputs := strings.Split(sc.Text(), " ")   // 半角スペース区切りでstring型として配列inputsに格納
		intInputs[0], _ = strconv.Atoi(inputs[0]) // string→intに型変換
		intInputs[1], _ = strconv.Atoi(inputs[1]) // string→intに型変換
		inputSliceMap[i] = intInputs
	}
	return n, inputSliceMap
}

// SingleColumnsBySlice  入力パターン
// N
// U1
// U2
// ...
// 0
// []int
func SingleColumnsBySlice() []int {
	sc := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ

	var inputSlice []int
	var input int
	for input > 0 { // 行数分繰り返す
		sc.Scan() // １行分の入力を取得する
		input, _ = strconv.Atoi(sc.Text())
		inputSlice = append(inputSlice, input)
	}
	return inputSlice
}

// GetGrid 入力
// H W
// B11 B12 B1w ...
// B21 B22 B2w ...
// Bh1 Bh2 Bhw ...
func GetGrid() (int, int, [][]int) {
	// 入力
	var H, W int
	fmt.Scan(&H, &W)

	A := make([][]int, H)
	for i := range A {
		A[i] = make([]int, W)
		for j := range A[i] {
			fmt.Scan(&A[i][j])
		}
	}
	return H, W, A
}
