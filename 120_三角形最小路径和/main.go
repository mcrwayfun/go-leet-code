package main

import (
	"fmt"
	"math"
)

// 递归算法
// time complexity: O(2^n)
// space complexity: O(1)
func minimumTotal(triangle [][]int) int {
	result := math.MaxInt64
	dfs(triangle, 0, 0, 0, &result)
	return result
}

func dfs(triangle [][]int, x, y, sum int, result *int) {
	n := len(triangle)
	if x == n {
		*result = min(*result, sum)
		return
	}

	dfs(triangle, x+1, y, sum+triangle[x][y], result)
	dfs(triangle, x+1, y+1, sum+triangle[x][y], result)
}

// 分治算法
// time complexity: O(n^2)
// space complexity: O(n^2)
func minimumTotalV2(triangle [][]int) int {
	if len(triangle) == 0 {
		return -1
	}
	if len(triangle[0]) == 0 {
		return -1
	}

	hash := initHash(triangle, math.MaxInt64)

	return divideConquer(triangle, 0, 0, hash)
}

func divideConquer(triangle [][]int, x, y int, hash [][]int) int {
	n := len(triangle)
	if x == n {
		return 0
	}

	if hash[x][y] != math.MaxInt64 {
		return hash[x][y]
	}

	hash[x][y] = triangle[x][y] + min(
		divideConquer(triangle, x+1, y, hash),
		divideConquer(triangle, x+1, y+1, hash))
	return hash[x][y]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func initHash(triangle [][]int, num int) [][]int {
	hash := make([][]int, len(triangle))
	for i := range hash {
		hash[i] = make([]int, len(triangle[i]))
	}

	m := len(triangle)
	for i := 0; i < m; i++ {
		n := len(triangle[i])
		for j := 0; j < n; j++ {
			hash[i][j] = num
		}
	}
	return hash
}

func main() {
	nums := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	result := minimumTotal(nums)
	fmt.Println(result)

	nums = [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	result = minimumTotalV2(nums)
	fmt.Println(result)
}
