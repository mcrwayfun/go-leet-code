package _40_search_a_2D_matrixII

/**
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

链接：https://leetcode-cn.com/problems/search-a-2d-matrix-ii
time complexity:O(m)+O(n)
space complexity:O(1)
*/
func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 ||
		len(matrix[0]) == 0 {
		return false
	}

	// 从左下角开始
	m, n := len(matrix)-1, 0
	for n <= len(matrix[0])-1 && m >= 0 {
		if matrix[m][n] == target {
			return true
		} else if matrix[m][n] < target {
			n++
		} else {
			m--
		}
	}

	return false
}
