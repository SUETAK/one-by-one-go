package main

import "fmt"

func integerInput() int {
	var s int
	fmt.Scanf("%d", &s)

	return s
}

func main() {
	s := integerInput()
	fmt.Println(integerDivisionReturns(s))
}

func integerDivisionReturns(s int) int {
	if s%10 == 0 {
		return s / 10
	}
	f := s / 10
	if f <= 0 && s < 0 {
		return f
	}
	return (s / 10) + 1
}
