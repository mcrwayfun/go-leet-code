package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	var m = len(matrix)
	var n = len(matrix[0])
	var low = 0
	var high = m*n - 1
	for low <= high {
		mid := low + (high-low)/2
		row := mid/n
		line := mid%n
		if matrix[row][line] == target {
			return true
		} else if matrix[row][line] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
