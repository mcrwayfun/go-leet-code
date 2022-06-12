package main

/**
编写一个高效的算法来判断m x n矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。

示例 1：
输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
输出：true

示例 2：
输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
输出：false

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/search-a-2d-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

题目：func searchMatrix(matrix [][]int, target int) bool {}
*/

/**
解题：由于原数组是一个z形增长的二维数组，可以将其看作一个递增的一维数组。
它们两个之间的下标转换有：m为行数，n为列数
1：row = index/n
2：column = index%n

time complexity: O(m*n)
space complexity: O(1)
*/
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
		r := mid / n
		c := mid % n
		if matrix[r][c] < target {
			low = mid + 1
		} else if matrix[r][c] > target {
			high = mid - 1
		} else {
			return true
		}
	}
	return false
}
