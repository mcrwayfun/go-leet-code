package main

// time complexity: O(n)
// space complexity: O(1)
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var buy = prices[0]
	var maxProfit int
	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {// 逢低买入
			buy = prices[i]
		} else {
			maxProfit = max(maxProfit, prices[i] - buy) // 计算最大的利润
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


