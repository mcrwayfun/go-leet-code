package main

// time complexity: O(n^2)
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
		d[i] = d[i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		d[0] += grid[i][0]
		for j := 1; j < n; j++ {
			d[j] = min(d[j], d[j-1]) + grid[i][j]
		}
	}
	return d[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
