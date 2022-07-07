package main

import "fmt"

/**
给定一个包含非负整数的 mxn网格grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。


示例 1：
输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
示例 2：

输入：grid = [[1,2,3],[4,5,6]]
输出：12

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/minimum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func minPathSum(grid [][]int) int {}
*/

/**
解题思路：
假设存储最优解的数据为d，那么当前坐标的最优解为d(i,j), 结果最优解为d(m-1,n-1)
1：对于首行，只能向右移动，有：d(0,j)=d(0,j-1)+a(0,j)
2：对于首列，只能向下移动，有：d(i,0)=d(i-1,0)+a(i,0)
3：对于其他元素，值可能来自于上左，则有：d(i,j)=min(d(i-1,j), d(i,j-1)) + a(i,j)

time complexity: O(m*n)
space complexity: O(m*n)
*/
func minPathSum(grid [][]int) int {
	var m = len(grid)
	var n = len(grid[0])

	var d = make([][]int, m)
	for i := 0; i < m; i++ {
		d[i] = make([]int, n)
	}

	d[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		d[i][0] = d[i-1][0] + grid[i][0]
	}

	for j := 1; j < n; j++ {
		d[0][j] = d[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			d[i][j] = min(d[i-1][j], d[i][j-1]) + grid[i][j]
		}
	}

	return d[m-1][n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/**
可以使用一维数组来存储最优路径解d[j]
1：对于首行有：d[j]=d[j-1]+a[0,j]
2：对于第二行开始有：
	2-1：初始化行首位元素，由上一行的首位+当前元素求值，d[0]=d[0]+a[i][0]
	2-2：d[j]=min(d[j],d[j-1])+a[i][j]

time complexity: O(m*n)
space complexity: O(n)
*/
func minPathSumWith1DArray(grid [][]int) int {
	var m = len(grid)
	var n = len(grid[0])

	d := make([]int, n)
	d[0] = grid[0][0]
	for j := 1; j < n; j++ {
		d[j] = d[j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		d[0] += grid[i][0]
		for j := 1; j < n; j++ {
			d[j] = min(d[j-1], d[j]) + grid[i][j]
		}
	}

	return d[n-1]
}

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	pathSum := minPathSum(grid)
	fmt.Println(pathSum)

	grid = [][]int{{1, 3, 1}, {4, 5, 6}}
	pathSum = minPathSum(grid)
	fmt.Println(pathSum)

	grid = [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	pathSum = minPathSumWith1DArray(grid)
	fmt.Println(pathSum)

	grid = [][]int{{1, 3, 1}, {4, 5, 6}}
	pathSum = minPathSumWith1DArray(grid)
	fmt.Println(pathSum)
}
