package main

import "fmt"

// 下記を解いた
// https://www.hackerrank.com/challenges/the-birthday-bar/submissions/code/357178798
func main() {
	// slice の部分取得のやり方
	//         0, 1, 2, 3, 4
	s := []int{1, 2, 3, 4, 5}

	choice := s[1:3] // index=1 ~ index=3 の一つ前までを取得するので2, 3

	fmt.Println(choice)
}
