package main

import "fmt"

// time complexity: O(n^2)
// space complexity: O(n^2)
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return -1
	}

	if len(grid[0]) == 0 {
		return -1
	}

	m := len(grid)
	n := len(grid[0])

	// init sum
	sum := initSum(m, n)
	sum[0][0] = grid[0][0]

	// init sum first row
	for i := 1; i < m; i++ {
		sum[i][0] = sum[i-1][0] + grid[i][0]
	}
	// init sum first column
	for i := 1; i < n; i++ {
		sum[0][i] = sum[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			sum[i][j] = min(sum[i-1][j], sum[i][j-1]) + grid[i][j]
		}
	}

	return sum[m-1][n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func initSum(m, n int) [][]int {
	sum := make([][]int, m)
	for i := range sum {
		sum[i] = make([]int, n)
	}
	return sum
}

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	pathSum := minPathSum(grid)
	fmt.Println(pathSum)

	grid = [][]int{{1, 3, 1}, {4, 5, 6}}
	pathSum = minPathSum(grid)
	fmt.Println(pathSum)
}
