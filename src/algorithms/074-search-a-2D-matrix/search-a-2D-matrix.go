package _74_search_a_2D_matrix

/**
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。

链接：https://leetcode-cn.com/problems/search-a-2d-matrix
time complexity:lg(mn)
space complexity:O(1)
*/
func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 ||
		len(matrix[0]) == 0 {
		return false
	}

	startRow, endRow := 0, len(matrix)-1
	endCol := len(matrix[0]) - 1
	for startRow+1 < endRow {
		mid := startRow + (endRow-startRow)/2
		if matrix[mid][endCol] == target {
			return true
		} else if matrix[mid][endCol] < target {
			startRow = mid
		} else {
			endRow = mid
		}
	}
	row := 0
	if matrix[startRow][endCol] >= target {
		row = startRow
	} else if matrix[endRow][endCol] >= target {
		row = endRow
	} else {
		return false
	}

	startCol := 0
	for startCol+1 < endCol {
		mid := startCol + (endCol-startCol)/2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			startCol = mid
		} else {
			endCol = mid
		}
	}

	if matrix[row][startCol] == target || matrix[row][endCol] == target {
		return true
	}

	return false
}
