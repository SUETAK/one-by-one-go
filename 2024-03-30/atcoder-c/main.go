package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	_, a, b, ps := notDecideInputAmountSeparatedBySpace()
	result := isAllHoliday(a, b, ps)
	if result {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func isAllHoliday(a, b int, d []int) bool {
	amariMap := map[int]int{}
	for i := 0; i < len(d)-1; i++ {
		amariMap[d[i]] = (a + b) % d[i]
	}
	e := []int{}
	for _, v := range amariMap {
		e = append(e, v)
	}
	sort.Ints(e)
	for i := 0; i < len(e)-2; i++ {
		if (e[i+1]-e[i])%(a+b) > b {
			return true
		} else {
			return false
		}
	}

	return true
}

func notDecideInputAmountSeparatedBySpace() (int, int, int, []int) {
	scanner := bufio.NewScanner(os.Stdin)
	// 1行目の読み込み (nとk)
	scanner.Scan()
	nk := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(nk[0])
	a, _ := strconv.Atoi(nk[1])
	b, _ := strconv.Atoi(nk[2])

	// 2行目の読み込み (ps)
	scanner.Scan()
	psStr := strings.Split(scanner.Text(), " ")
	ps := make([]int, n)
	for i, pStr := range psStr {
		ps[i], _ = strconv.Atoi(pStr)
	}
	return n, a, b, ps
}
