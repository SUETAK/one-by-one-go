package main

func lengthOfLongestSubstring_answer(s string) int {
	// A map to store the last index of each character
	lastIndex := make(map[byte]int)
	maxLength := 0
	start := 0 // The starting index of the current window

	for i := 0; i < len(s); i++ {
		if idx, found := lastIndex[s[i]]; found && idx >= start {
			start = idx + 1
		}
		lastIndex[s[i]] = i
		if length := i - start + 1; length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}
