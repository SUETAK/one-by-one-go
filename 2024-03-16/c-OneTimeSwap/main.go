package main

import (
	"fmt"
)

func stringInput() string {
	var s string
	fmt.Scanf("%s", &s)

	return s
}

func main() {
	s := stringInput()

	fmt.Println(oneTimeSwap(s))
}

func oneTimeSwap(s string) string {
	result := map[string]bool{}

	for i := 0; i < len(s)-1; i++ {
		head := s[i]
		for j := i+1; j < len(s)-1; j++ {
			for k := 0; k < ; k++ {

			}
		}
	}
}
