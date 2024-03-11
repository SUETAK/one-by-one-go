package main

import "fmt"

func main() {
	var a int
	fmt.Scanf("%d", &a)

	answerSlice := answerCorrectBrackets(a)

	for _, an := range answerSlice {
		fmt.Println(an)
	}
}
