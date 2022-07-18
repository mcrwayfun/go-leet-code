package _1

// time complexity: O(n)
// space complexity: O(1)
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var maxProfit int
	var buy = prices[0]
	for i := 1; i < len(prices); i++ {
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
