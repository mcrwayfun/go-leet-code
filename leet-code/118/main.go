package main

import "fmt"

/**
给定一个非负整数numRows，生成「杨辉三角」的前numRows行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。

示例 1:
输入: numRows = 5
输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

示例2:
输入: numRows = 1
输出: [[1]]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/pascals-triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func generate(numRows int) [][]int {}
*/

/**
1
1 1
1 2 1
1 3 3 1
1 4 6 4 1

从题目可以得知，i从下标0开始，第i行共有i+1个元素
三角形左右两边的元素都是1：a(i,0)=a(i,i)=1
三角形中间的元素：a(i,j)=a(i-1, j-1)+a(i-1,j)

time complexity: O(n^2)
space complexity: O(1)
*/
func generate(numRows int) [][]int {
	if numRows < 1 {
		return [][]int{}
	}

	var ret = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		var rows = make([]int, i+1)
		rows[0] = 1
		rows[i] = 1
		for j := 1; j < i; j++ {
			rows[j] = ret[i-1][j-1] + ret[i-1][j]
		}
		ret[i] = rows
	}
	return ret
}

func main() {
	fmt.Println(generate(5))
}