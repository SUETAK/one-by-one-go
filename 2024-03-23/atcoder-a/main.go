package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	nums := notDecideInputAmountSeparatedBySpace()
	results := calc(nums)
	fmt.Println(strings.Join(results, " "))

}

func calc(nums []int) []string {
	results := make([]string, len(nums)-1, len(nums)-1)
	for i := 0; i < len(nums)-1; i++ {
		results[i] = strconv.Itoa(nums[i] * nums[i+1])
	}
	return results
}
