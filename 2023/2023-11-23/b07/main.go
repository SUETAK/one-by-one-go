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
	sc.Scan()

	closedTime, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	time := make([]int, closedTime)
	for i := 0; i < n; i++ {
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")
		L, _ := strconv.Atoi(inputs[0])
		R, _ := strconv.Atoi(inputs[1])

		time[L]++
		if R == closedTime {
			continue
		}
		time[R]--
		//fmt.Println(time)
	}

	answer := make([]int, closedTime)
	answer[0] = time[0]
	for t := 1; t < closedTime; t++ {
		answer[t] = answer[t-1] + time[t]
	}

	for _, v := range answer {
		fmt.Println(v)
	}
}
