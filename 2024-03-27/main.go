package main

func main() {}

func maxProfit(prices []int) int {
	maxP := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		price := prices[i]
		if minPrice > price {
			minPrice = price
		} else {
			profit := price - minPrice
			if profit > maxP {
				maxP = profit
			}
		}
	}

	return maxP
}
