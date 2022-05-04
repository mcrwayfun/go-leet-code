package main

/**
编写一个高效的算法来搜索mxn矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-a-2d-matrix-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func searchMatrix(matrix [][]int, target int) bool
*/

/**
解题思路：这道题给出的条件有，
1：每一行从左到右是递增的
2：每一列从上到下是递增的
3：不能够保证下一行的首个数字比上一行每个数字都大，列同理
由此可以得出的结论是，
1：从矩阵的右上角开始，同一行所有数字都比它小，同一列所有数字都比它大
2：从矩阵的左下角开始，同一行所有数字都比它大，同一列所有数字都比它小
所以解法可以从右上角或者左下角开始，现我们从右上角开始，坐标为(m,n)
1：如果target比matrix[m,n]小，则同一列的都不可能，则列需要减少，即n--
2：如果target比matrix[m,n]大，则同一行的都不可能，则行需要增加，即m++
3：如果target==matrix[m,n]，return m,n

time complexity: O(m+n)
space complexity: O(1)
*/
func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || matrix[0] == nil {
		return false
	}

	m := len(matrix)
	n := len(matrix[0])

	i := 0
	j := n - 1
	for i < m && j >= 0 {
		if target < matrix[i][j] {
			j--
		} else if target > matrix[i][j] {
			i++
		} else if target == matrix[i][j] {
			return true
		}
	}
	return false
}
