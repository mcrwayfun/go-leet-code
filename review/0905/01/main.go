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

	var prev2 = cost[0]
	var prev1 = cost[1]
	for i := 2; i < len(cost); i++ {
		cur := min(prev1, prev2) + cost[i]
		prev2 = prev1
		prev1 = cur
	}
	return min(prev1, prev2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
