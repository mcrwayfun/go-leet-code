package main

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	var m = len(matrix)
	var n = len(matrix[0])

	var rows = make([]bool, m)
	var columns = make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				columns[j] = true
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rows[i] || columns[j] {
				matrix[i][j] = 0
			}
		}
	}

	return
}
