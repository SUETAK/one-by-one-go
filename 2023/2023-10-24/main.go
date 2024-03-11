package main

func main() {
	// words = ["duel", "speed", "duel", "cars"]
	// queries = ["spede", "deul"]

	// queries がwords のアナグラムになっているかを考える

	//for i, i2 := range words {
	//
	//}

}

func isAnagram(word, query string) bool {
	if len(word) != len(query) {
		return false
	}
	// word = "duel"
	// query = "deul"
	// wordMap = map[string]int{"d": 1, "u": 1, "e": 1, "l": 1}
	// queryMap = map[string]int{"d": 1, "e": 1, "u": 1, "l": 1}
	wordMap := make(map[string]int)
	queryMap := make(map[string]int)
	for _, char := range word {
		wordMap[string(char)]++
	}
	for _, char := range query {
		queryMap[string(char)]++
	}
	for key, value := range wordMap {
		if queryMap[key] != value {
			return false
		}
	}
	return true

}
