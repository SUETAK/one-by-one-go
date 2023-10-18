package main

import "fmt"

func main() {
	a, b, c := decideInputAmountNotSeparated()

	ans := countPutOnBallCells(a, b, c)

	fmt.Println(ans)
}
