package atcoder

import (
	"bufio"
	"fmt"
	"os"
)

func mainC() {
	_, ps, bs, cs, xs := decideInputCountAndNotDecideInputAmountSeparatedBySpace()
	aPlusBPlusC(ps, bs, cs, xs)
}

func aPlusBPlusC(as, bs, cs, xs []int) []bool {
	// xi の値をas, bs, cs のうちどれかの要素で表現できるかどうかを判定する
	// 100*100*100 を1回行って、作成した値をset に入れる。答えとの称号をset を使って判定する
	var valueMap = make(map[int]bool)

	// 絶対に1つは選ぶ必要がある
	for _, av := range as {
		for _, bv := range bs {
			for _, cv := range cs {
				valueMap[av+bv+cv] = true
			}
		}
	}

	var result []bool
	for _, x := range xs {
		if valueMap[x] {
			//result = append(result, true)
			fmt.Println("Yes")
		} else {
			//result = append(result, false)
			fmt.Println("No")
		}
	}
	return result
}

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
