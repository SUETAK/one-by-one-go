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
			// Hashmap to store characters of the current substring
			// Needed to check if all the characters are unique or if it contains duplicates
			//　現在のsubstring の文字を保存するためのハッシュマップ
			// すべての文字がユニークかどうか、または重複が含まれているかを確認するために必要
			characterMap := make(map[uint8]bool)

			// True if substring contains all unique characters
			// False if it contains duplicates
			// substring がすべてユニークな文字を含む場合はTrue
			// 重複が含まれている場合はFalse
			isUniqueSubstring := true

			// Loop to check if the current substring contains unique characters or has duplicates
			// 現在のsubstring がユニークな文字を含むか、重複が含まれているかを確認するためのループ
			for i := start; i <= end; i++ {
				// If the character is not present in the substring
				// Add it to the hashmap and continue to the next character in the substring
				// 文字がsubstring に存在しない場合
				// ハッシュマップに追加して、次の文字に進む
				if _, isPresent := characterMap[s[i]]; !isPresent {
					characterMap[s[i]] = true
					continue // continue to next character in the substring
				}

				// substring contains duplicates
				// 重複が含まれている
				isUniqueSubstring = false
				break
			}

			// If the current substring has all unique characters
			// Update result if current substring length is greater than result so far
			// 現在のsubstring がすべてユニークな文字を含む場合
			// 現在のsubstring の長さがこれまでの結果よりも大きい場合は、結果を更新する
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
