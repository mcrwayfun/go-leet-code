package main

import "fmt"

// time complexity: O(n^2)
// space complexity: O(n^2)
func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	fn := initFn(m, n)
	// 到第一列每个节点的路径为1
	for i := 0; i < m; i++ {
		fn[i][0] = 1
	}
	// 到第一行每个节点的路径为1
	for i := 0; i < n; i++ {
		fn[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			fn[i][j] = fn[i-1][j] + fn[i][j-1]
		}
	}

	return fn[m-1][n-1]
}

func initFn(m, n int) [][]int {
	fn := make([][]int, m)
	for i := range fn {
		fn[i] = make([]int, n)
	}
	return fn
}

func main() {
	paths := uniquePaths(3, 7)
	fmt.Println(paths)

	paths = uniquePaths(3, 2)
	fmt.Println(paths)

	paths = uniquePaths(7, 3)
	fmt.Println(paths)
}
