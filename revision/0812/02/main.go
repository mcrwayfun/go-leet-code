package main

// time complexity: O(lgn)
// space complexity: O(1)
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	var m = len(matrix)
	var n = len(matrix[0])
	var start int
	var end = m*n - 1
	for start <= end {
		mid := start + (end-start)/2
		x := mid/n
		y := mid%n
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}
