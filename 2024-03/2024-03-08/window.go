package main

func lengthOfLongestSubstringWindow(s string) int {
	// Length of the given input string
	// input string の長さ
	n := len(s)

	// Length of longest substring
	// 最長のsubstring の長さ
	var result int

	// Map to store visited characters along with their index
	// 文字とそのインデックスを保存するためのマップ
	charIndexMap := make(map[uint8]int)

	// start indicates the start of current substring
	// 現在のsubstring の開始を示す
	var start int

	// Iterate through the string and slide the window each time you find a duplicate
	// end indicates the end of current substring
	// 文字列を反復処理し、重複を見つけるたびにウィンドウをスライドする
	// end は現在のsubstring の終わりを示す
	for end := 0; end < n; end++ {

		// Check if the current character is a duplicate
		// 現在の文字が重複しているかどうかを確認する
		// s="abcabcd" なら s[end] は a から始まり、a が重複しているので、if 文に入る
		duplicateIndex, isDuplicate := charIndexMap[s[end]]
		if isDuplicate {
			// Update the result for the substring in the current window before we found duplicate character
			// 重複文字を見つける前の現在のウィンドウのsubstring の結果を更新する
			result = max(result, end-start)

			// Remove all characters before the new
			// 新しい重複文字の前のすべての文字を削除する
			for i := start; i <= duplicateIndex; i++ {
				delete(charIndexMap, s[i])
			}

			// Slide the window since we have found a duplicate character
			// 重複文字を見つけたので、ウィンドウをスライドする
			start = duplicateIndex + 1
		}
		// Add the current character to hashmap
		// 現在の文字をハッシュマップに追加する
		charIndexMap[s[end]] = end
	}
	// Update the result for last window
	// For a input like "abc" the execution will not enter the above if statement for checking duplicates
	// 最後のウィンドウの結果を更新する
	// "abc" のような入力の場合、重複をチェックするための上記のif 文には入らない
	result = max(result, n-start)

	return result
}
