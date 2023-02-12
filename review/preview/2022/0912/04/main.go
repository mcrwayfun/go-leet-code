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
		if obstacleGrid[i][0] == 1 {
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

func uniquePathsWithObstaclesV2(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	var m = len(obstacleGrid)
	var n = len(obstacleGrid[0])
	var d = make([][]int, m)
	for i := range d {
		d[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		d[i][0] = 1
	}

	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		d[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				d[i][j] = 0
			} else {
				d[i][j] = d[i-1][j] + d[i][j-1]
			}
		}
	}

	return d[m-1][n-1]
}
