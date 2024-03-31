package main

import "strings"

func main() {}

func convert(s string, numRows int) string {
	// 行数が１ならそのまま返す
	if numRows == 1 {
		return s
	}

	// s の長さ分のstrings を確保
	var b strings.Builder
	b.Grow(len(s))
	// 0行目の1文字目~2文字目のstep数
	step := numRows*2 - 2

	for row := 0; row < numRows; row++ {
		nextStep := 0
		// 0行目か、最終rowなら、次のstep 数は nomRows*2 - 2
		if row == 0 || row == numRows-1 {
			nextStep = step
		} else {
			nextStep = row * 2
		}
		for i := row; i < len(s); i += nextStep {
			b.WriteByte(s[i])
			if row != 0 && row != numRows-1 {
				nextStep = step - nextStep
			}
		}
	}

	return b.String()
}
