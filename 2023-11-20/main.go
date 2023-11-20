package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")

	sc.Scan()

	for sc.Scan() {
		input := sc.Text()
		numB, _ := strconv.Atoi(input)
		score := math.MaxInt
		for _, numA := range inputs {
			numA, _ := strconv.Atoi(numA)
			d := dissatisfaction(numA, numB)
			if d < score {
				score = d
			}
		}
		fmt.Println(score)
	}
}

func dissatisfaction(rateA, rateB int) int {
	return int(math.Abs(float64(rateA - rateB)))
}
