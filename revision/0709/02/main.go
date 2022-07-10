package main

func minPathSum(grid [][]int) int {
	var m = len(grid)
	var n = len(grid[0])

	var d = make([]int, n)
	d[0] = grid[0][0]
	for j:=1; j<n; j++ {
		d[j] = d[j-1] + grid[0][j]
	}

	for i:=1; i<m; i++ {
		d[0] += grid[i][0]
		for j:=1; j<n; j++ {
			d[j] = min(d[j], d[j-1]) + grid[i][j]
		}
	}

	return d[n-1]
}

//
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
