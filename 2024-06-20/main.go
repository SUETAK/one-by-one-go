package main

import (
	"fmt"
	se "one-by-one-go/2024-06-20/struct_embedding"
)

func main() {

	a := &se.A{B: se.B{Id: 1}, I: se.C{Name: "Struct C"}}

	fmt.Println(a.GetId())
	fmt.Println(a.B.GetId())

	fmt.Println(a.GetName())
	fmt.Println(a.I.GetName())
}
