package atcoder

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main2() {
	ps := SingleColumnsBySlice()
	result := delimiter(ps)
	for _, r := range result {
		fmt.Println(r)
	}
}

func delimiter(ps []int) []int {
	var result []int
	for i := 0; i < len(ps); i++ {
		backIndex := len(ps) - 1 - i
		result = append(result, ps[backIndex])
	}

	return result
}

func SingleColumnsBySlice() []int {
	sc := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ

	var inputSlice []int
	var input int
	for input > 0 { // 行数分繰り返す
		sc.Scan() // １行分の入力を取得する
		input, _ = strconv.Atoi(sc.Text())
		inputSlice = append(inputSlice, input)
	}
	return inputSlice
}
