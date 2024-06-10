package main

func main() {}

func isIsomorphic(s string, t string) bool {
	map1 := make(map[rune]rune) // Stores mapping of s to t
	map2 := make(map[rune]rune) // Stores mapping of t to s

	for i, sCh := range s {
		tCh := rune(t[i])

		sValue, existsSValue := map1[sCh]
		tValue, existsTValue := map2[tCh]

		// 両方にデータがないなら保存
		if !existsSValue && !existsTValue {
			// 保存する時はs文字のvalueをt文字にする。逆も然り
			map1[sCh] = tCh
			map2[tCh] = sCh
		} else if existsSValue && sValue != tCh {
			// s文字のvalueをt文字にしているので、もしs文字が保存されているvalue がt文字でない場合は同型じゃなくなる
			return false
		} else if existsTValue && tValue != sCh {
			return false
		}
	}
	return true
}
