package main

import (
	"fmt"
	"one-by-one-go/2023/2023-10-12/domain"
)

func main() {
	// Stackの型パラメータにstringを指定
	s := domain.TwinKataStack[string]{}

	s.Push("a")
	s.Push("b")

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	// Stackの型パラメータにanyを指定
	stack := domain.TwinKataStack[any]{}
	stack.Push(1)
	stack.Push("a")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	domain.Authors()

}
