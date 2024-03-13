package main

func main() {}

func longestPalindrome(s string) string {

	var answer string
	var tmp string
	for _, v := range s {
		if len(tmp) < 2 {
			tmp = tmp + string(v)
			if tmp[0] == tmp[1] {
				answer = tmp
			} else {
				answer = string(tmp[1])
			}
			continue
		}
		if len(answer) == 1 {
			answer = answer + string(v)
			continue
		}
		if len(answer) == 2 {

		}
		// answer が偶数の場合
		if len(answer)%2 == 0 {
			if string(v) == string(answer[len(answer)-1]) {
				answer = answer + string(v)
			}
		} else {

		}

	}

	return answer
}
