package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var buy = prices[0]
	var maxProfit int
	for i:=1; i<len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
		} else {
			maxProfit = max(maxProfit, prices[i]-buy)
		}
	}
	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
