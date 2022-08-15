package main

// time complexity: O(m*n)
// space complexity: O(n)
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	var m = len(grid)
	var n = len(grid[0])
	var d = make([]int, n)

	d[0] = grid[0][0]
	for i := 1; i < n; i++ {
		d[i] = grid[0][i] + d[i-1]
	}

	for i := 1; i < m; i++ {
		d[0] = grid[i][0]
		for j := 1; j < n; j++ {
			d[j] = max(d[j-1], d[j]) + grid[i][j]
		}
	}
	return d[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
