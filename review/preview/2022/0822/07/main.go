package main

// time complexity: O(m*n)
// space complexity: O(m*n)
func uniquePaths(m int, n int) int {
	if m < 1 || n < 1 {
		return 0
	}

	var d = make([][]int, m)
	for i := range d {
		d[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		d[i][0] = 1
	}
	for j := 0; j < n; j++ {
		d[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			d[i][j] = d[i-1][j] + d[i][j-1]
		}
	}

	return d[m-1][n-1]
}
