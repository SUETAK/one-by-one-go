package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i

}

func scanInput() (int, int, []int) {
	n, k := scanInt(), scanInt()
	a := []int{}
	for i := 0; i < n; i++ {
		a = append(a, scanInt())
	}
	return n, k, a
}

func main() {
	_, k, ary := scanInput()
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
