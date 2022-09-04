package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	var m = len(obstacleGrid)
	var n = len(obstacleGrid[0])
	var d = make([]int, n)
	if obstacleGrid[0][0] == 1 {
		d[0] = 0
	} else {
		d[0] = 1
	}

	for i := 0; i < m; i++ {
		if obstacleGrid[0][i] == 1 {
			d[0] = 0
		}

		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				d[j] = 0
			} else {
				d[j] = d[j] + d[j-1]
			}
		}
	}

	return d[n-1]
}
