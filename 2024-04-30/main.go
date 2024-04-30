package main

import "strings"

func main() {}

func fullJustify(words []string, maxWidth int) []string {
	wordCount := 0
	spaceCount := 0
	tmpWords := make([]string, 0)
	result := make([]string, 0)
	for i := 0; i < len(words); i++ {
		word := words[i]
		wl := len(word)

		wordCount += wl + spaceCount
		if wordCount >= maxWidth {
			result = append(result, createStrings(tmpWords, wordCount-wl-spaceCount, maxWidth))
			wordCount = 0
			spaceCount = 0
			tmpWords = make([]string, 0)
			continue
		} else {
			tmpWords = append(tmpWords, word)
		}
		spaceCount++
	}
	return result
}

func createStrings(words []string, wordsCount int, maxWidth int) string {
	// words の間には" " が入る
	if len(words) == 2 {
		return strings.Join(words, " ") + strings.Repeat(" ", maxWidth-wordsCount-1)
	}
	if len(words) == 3 {
		left := (maxWidth - wordsCount) / (len(words) - 1)
		right := maxWidth - wordsCount - left
		return words[0] + strings.Repeat(" ", left) + words[1] + strings.Repeat(" ", right) + words[2]
	}
	return ""
}
