package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// １行分の入力を取得する

	sc.Scan()
	d, _ := strconv.Atoi(sc.Text())
	darr := make([]int, d)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	for i := 0; i < n; i++ {
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")
		// l = 2
		l, _ := strconv.Atoi(inputs[0])
		// r = 4
		r, _ := strconv.Atoi(inputs[1])
		// [0, 0, 1, -1, 0]
		// [0, 1, 0, 0, -1]
		darr[l-1]++
		if r == d {
			continue
		}
		darr[r]--
		//fmt.Println(darr)
	}

	answer := make([]int, d)
	answer[0] = darr[0]
	for day := 1; day < d; day++ {
		//fmt.Printf("day-1=%d \n", answer[day-1])
		answer[day] = answer[day-1] + darr[day]
		//fmt.Println(answer)
	}

	for _, v := range answer {
		fmt.Println(v)
	}

}
