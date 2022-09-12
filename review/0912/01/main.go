package main

// time complexity: O(n^2)
// space complexity: O(n^2)
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	var m = len(matrix)
	var n = len(matrix[0])
	var rows = make([]bool, m)
	var lines = make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				lines[j] = true
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rows[i] || lines[j] {
				matrix[i][j] = 0
			}
		}
	}
}
