package main

import (
	"math"
	"strconv"
	"strings"
)

func answerCorrectBrackets(a int) []string {
	if a%2 == 1 {
		return []string{""}
	}
	// 組み合わせの数
	combinationCount := math.Pow(2, float64(a))

	var combinations []string
	// 組み合わせの数だけループ。後半の組み合わせは不要なので半分だけループ
	for i := 0; i < int(combinationCount)/2; i++ {

		// FormatInt関数を使って数値を2進数の文字列に変換
		binString := strconv.FormatInt(int64(i), 2)
		// aビットになるように0でパディング
		for len(binString) < a {
			binString = "0" + binString
		}
		combinations = append(combinations, binString)
	}

	var correctCombinations []string
	for _, combination := range combinations {
		if isValid(combination) {
			combination = strings.ReplaceAll(combination, "0", "(")
			combination = strings.ReplaceAll(combination, "1", ")")
			correctCombinations = append(correctCombinations, combination)
		}
	}
	return unique(correctCombinations)
}

func isValid(s string) bool {
	score := 0
	for _, char := range s {
		if char == '0' {
			score++
		} else if char == '1' {
			score--
		}
		if score < 0 {
			return false
		}
	}
	return score == 0
}

func unique(slice []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}
