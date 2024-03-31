package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()
	fmt.Println(getSubstrings(s))
}

func getSubstrings(s string) int {
	n := len(s)
	if n == 1 {
		return 1
	}
	result := 0
	charMap := make(map[string]bool)
	for start := 0; start <= n-1; start++ {
		for end := start + 1; end <= n; end++ {
			value := string(s[start:end])
			if _, isPresent := charMap[value]; !isPresent {
				if value != "" {
					charMap[value] = true
					result++
				}
			}
		}

	}
	return result
}
