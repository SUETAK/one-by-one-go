package main

import "fmt"

func main() {
	n, numbers := decideInputCountAndNotDecideInputAmountSeparatedBySpace()

	ans := countMaxControl(n, numbers)

	fmt.Println(ans)
}
