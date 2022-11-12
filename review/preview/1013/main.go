package main

// time complexity: O(n)
// space complexity: O(1)
func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}

	if len(cost) == 1 {
		return cost[0]
	}

	var first = cost[0]
	var second = cost[1]
	for i := 2; i < len(cost); i++ {
		cur := min(first, second) + cost[i]

		first = second
		second = cur
	}
	return min(first, second)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
