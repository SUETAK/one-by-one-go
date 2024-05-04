package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	sc := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ
	sc.Scan()                        // １行分の入力を取得する
	S := sc.Text()
	sc.Scan() // １行分の入力を取得する
	T := sc.Text()

	cti := correctTypeIndex(S, T)
	fmt.Println(strings.Join(cti, " "))
}

func correctTypeIndex(s, t string) []string {
	result := make([]string, 0)
	k := 0
	for i := 0; i < len(s); i++ {
		sv := s[i]
		for k < len(t) {
			if sv == t[k] {
				result = append(result, strconv.Itoa(k+1))
				k = k + 1
				break
			}
			k++
		}
	}
	return result
}
