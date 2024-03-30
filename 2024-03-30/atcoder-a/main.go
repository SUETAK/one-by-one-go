package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	_, k, ps := notDecideInputAmountSeparatedBySpace()
	result := divideKValue(k, ps)
	fmt.Println(strings.Join(result, " "))
}

func divideKValue(k int, ps []int) []string {
	var result []string
	for i := 0; i < len(ps); i++ {
		if ps[i]%k == 0 {
			result = append(result, strconv.Itoa(ps[i]/k))
		}
	}
	return result
}

func notDecideInputAmountSeparatedBySpace() (int, int, []int) {
	scanner := bufio.NewScanner(os.Stdin)
	// 1行目の読み込み (nとk)
	scanner.Scan()
	nk := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(nk[0])
	k, _ := strconv.Atoi(nk[1])

	// 2行目の読み込み (ps)
	scanner.Scan()
	psStr := strings.Split(scanner.Text(), " ")
	ps := make([]int, n)
	for i, pStr := range psStr {
		ps[i], _ = strconv.Atoi(pStr)
	}
	return n, k, ps
}
