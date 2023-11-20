package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	nextInt := func() int {
		scanner.Scan()
		i, _ := strconv.Atoi(scanner.Text())
		return i
	}

	// 無限大を表す値
	const INF = 1 << 60

	// 入力
	N := nextInt()

	A := make([]int, N)
	for i := range A {
		A[i] = nextInt()
	}

	// ソート
	sort.Ints(A)

	// クエリに答える
	Q := nextInt()
	for i := 0; i < Q; i++ {
		B := nextInt()

		// A[j] >= B となる最小の j を求める
		j := sort.Search(len(A), func(i int) bool { return A[i] >= B })

		// A[j] と A[j-1] を比較する
		res := INF
		if j < N {
			res = min(res, abs(B-A[j]))
		}
		if j > 0 {
			res = min(res, abs(B-A[j-1]))
		}

		fmt.Println(res)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
