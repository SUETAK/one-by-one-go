package main

func main() {}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	// 利益、売り日、買い日
	profit := 0
	buy := 0
	sell := 0

	// 1から始める
	for i := 1; i < len(prices); i++ {
		currentValue := prices[i]
		previousValue := prices[i-1]
		switch {
		// 現在の価格が、前日の価格以上であれば、売り日を更新
		case currentValue >= previousValue:
			sell = i
		// 現在の価格が前日の価格より低ければ、売り日の価格-買い日の価格を計算し、利益に追加する
		case currentValue < previousValue:
			profit += prices[sell] - prices[buy]
			// 売ったことになるため、買い日と売り日をリセットする
			buy = i
			sell = i
		}
	}
	profit += prices[sell] - prices[buy]

	return profit
}
