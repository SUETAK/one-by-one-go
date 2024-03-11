package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func correctBrackets(a int) map[int]string {
	if a%2 == 1 {
		return map[int]string{0: ""}
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

	var correctCombinations map[int]string = make(map[int]string)
	for i, combination := range combinations {
		// 先頭と末尾が( ( の組み合わせはNGなので除外
		if string([]rune(combination)[0]) == string([]rune(combination)[len(combination)-1]) {
			continue
		}

		// 0と1の数が同じでない組み合わせはNGなので除外
		zeroCount, oneCount := countZerosAndOnes(combination)
		if zeroCount != oneCount {
			continue
		}

		// 以降は ( .... ) のような組み合わせだけが残る
		// combination をa/2 ビットずつ分割する
		// ((())) なら front:((( と back:))) に分割する
		front := combination[:a/2]
		back := combination[a/2:]
		// front, back をそれぞれソート
		sortedFront := sortString(front)
		sortedBack := sortString(back)
		// front, back を結合
		combination = sortedFront + sortedBack

		combination = strings.ReplaceAll(combination, "0", "(")
		combination = strings.ReplaceAll(combination, "1", ")")
		correctCombinations[i] = combination
	}
	return correctCombinations
}

func sortString(s string) string {
	slice := strings.Split(s, "")
	sort.Strings(slice)
	return strings.Join(slice, "")
}

func countZerosAndOnes(s string) (int, int) {
	var zeroCount, oneCount int
	for _, char := range s {
		if char == '0' {
			zeroCount++
		} else if char == '1' {
			oneCount++
		}
	}
	return zeroCount, oneCount
}
