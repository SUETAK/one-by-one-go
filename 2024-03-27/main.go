package main

func main() {}

func maxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		v := prices[i]
		for j := i + 1; j < len(prices); j++ {
			v2 := prices[j]
			if v2-v < 0 {
				continue
			}
			if profit < v2-v {
				profit = v2 - v
				continue
			}
		}
	}

	return profit
}
