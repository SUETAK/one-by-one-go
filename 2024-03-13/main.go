package main

func main() {}

func longestPalindrome(s string) string {
	// "a" や "ac" の時のことを考える
	answer := string(s[0])
	// 奇数の回文検索
	for i := 0; i < len(s); i++ {
		ahead := i - 1
		back := i + 1

		for ahead >= 0 && back < len(s) {
			if s[ahead] == s[back] {
				if len(s[ahead:back+1]) > len(answer) {
					answer = s[ahead : back+1]
				}
			} else {
				// 同じ文字でない時点で回文ではなくなるのでbreak
				break
			}
			ahead--
			back++
		}
	}

	for i := 0; i < len(s); i++ {
		ahead := i
		back := i + 1

		for ahead >= 0 && back < len(s) {
			if s[ahead] == s[back] {
				if len(s[ahead:back+1]) > len(answer) {
					answer = s[ahead : back+1]
				}
			} else {
				break
			}
			ahead--
			back++
		}
	}
	return answer
}
