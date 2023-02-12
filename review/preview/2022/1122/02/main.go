package main

// time complexity: O(n^2)
// space complexity: O(1)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	var m = len(obstacleGrid)
	var n = len(obstacleGrid[0])

	// init the distance 2D arrays
	var d = make([][]int, m)
	for i := 0; i < m; i++ {
		d[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		d[i][0] = 1
	}

	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		d[0][i] = 1
	}

	// get the result
	for x := 1; x < m; x++ {
		for y := 1; y < n; y++ {
			if obstacleGrid[x][y] == 1 {
				d[x][y] = 0
			} else {
				d[x][y] = d[x-1][y] + d[x][y-1]
			}
		}
	}

	return d[m-1][n-1]
}
