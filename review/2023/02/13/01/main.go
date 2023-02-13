package main

import "math"

// time complexity: O(n)
// space complexity: O(1)
func maxProfit(prices []int) int {
	var bestProfit int
	var price = math.MaxInt64
	for i := 0; i < len(prices); i++ {
		if prices[i] < price {
			price = prices[i]
		} else {
			bestProfit = max(bestProfit, prices[i]-price)
		}
	}
	return bestProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
