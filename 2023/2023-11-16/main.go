package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// input から標準入力を受け取る
func main() {
	var n int
	fmt.Scanf("%d", &n)

	phoneNumberMap := make(map[string]string)

	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")
		phoneNumberMap[inputs[0]] = inputs[1]
	}

	// 入力を全て読み込むまでfor loop を回す
	for sc.Scan() {
		if v, ok := phoneNumberMap[sc.Text()]; ok {
			fmt.Println(v)
		} else {
			fmt.Println("Not found")
		}
	}
}
