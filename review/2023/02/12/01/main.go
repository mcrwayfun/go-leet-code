package main

// time complexity: O(n^2)
// space complexity: O(n^2)
func minPathSum(grid [][]int) int {
	var m = len(grid)
	var n = len(grid[0])
	var dp = make([][]int, m)

	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 初始化 dp[i][0] 和 d[0][j]
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
