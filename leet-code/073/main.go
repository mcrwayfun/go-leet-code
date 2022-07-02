package main

/**
给定一个m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。


示例 1：
输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
输出：[[1,0,1],[0,0,0],[1,0,1]]

示例 2：
输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/set-matrix-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func setZeroes(matrix [][]int){}
 */

// time complexity: O(m*n)
// space complexity: O(m+n)
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
}
