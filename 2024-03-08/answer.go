package main

// 全探索
func lengthOfLongestSubstringAnswer(s string) int {
	n := len(s)

	// １文字だけなら常に固有
	if n == 1 {
		return 1
	}

	var result int
	for start := 0; start < n-1; start++ {
		for end := start; end < n; end++ {
			characterMap := make(map[uint8]bool)
			isUniqueSubstring := true
			for i := start; i <= end; i++ {
				if _, isPresent := characterMap[s[i]]; !isPresent {
					characterMap[s[i]] = true
					continue // continue to next character in the substring
				}
				isUniqueSubstring = false
				break
			}
			if isUniqueSubstring {
				result = max(result, end-start+1)
			}
			if result == 62 {
				break
			}
		}
	}
	return result
}

// max returns the max of num1 and num2
func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
