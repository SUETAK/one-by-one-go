package main

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	length := 0
	slice := strings.Split(s, "")
	tmp := make(map[string]string)

	// 文字数をfor文で回す
	for i, v := range slice {
		tmp[v] = v
		if length == 62 {
			break
		}
		// i の文字以降を回す
		for _, v2 := range slice[i+1:] {
			// map あるならlength を保管。tmpを初期化して
			if _, ok := tmp[v2]; ok {
				if len(tmp) > length {
					length = len(tmp)
				}
				tmp = make(map[string]string)
				tmp[v2] = v2
				continue
			}
			tmp[v2] = v2
		}
		if len(tmp) > length {
			length = len(tmp)
		}
		tmp = make(map[string]string)
	}

	if len(tmp) > length {
		length = len(tmp)
	}

	return length
}
