package main

import "fmt"

// time complexity: O(n^2)
// space complexity: O(n^2)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}

	if len(obstacleGrid[0]) == 0 {
		return 0
	}

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	fn := initFn(m, n)

	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] != 1 {
			fn[i][0] = 1
		} else {
			break
		}
	}

	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] != 1 {
			fn[0][i] = 1
		} else {
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				fn[i][j] = fn[i-1][j] + fn[i][j-1]
			} else {
				fn[i][j] = 0
			}
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
	nums := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	pathSum := uniquePathsWithObstacles(nums)
	fmt.Println(pathSum)

	nums = [][]int{{1, 0}}
	pathSum = uniquePathsWithObstacles(nums)
	fmt.Println(pathSum)

	nums = [][]int{{0, 0}, {1, 1}, {0, 0}}
	pathSum = uniquePathsWithObstacles(nums)
	fmt.Println(pathSum)
}
