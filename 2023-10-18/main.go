package main

import "fmt"

func main() {
	a, b := twoInputAmountSeparatedBySpace()

	ans := (a * b) % 2
	if ans == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
