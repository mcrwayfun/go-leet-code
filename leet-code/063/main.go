package main

/**
一个机器人位于一个m x n网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

网格中的障碍物和空位置分别用 1 和 0 来表示。

示例 1：
输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
输出：2
解释：3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右
示例 2：

输入：obstacleGrid = [[0,1],[0,0]]
输出：1
来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/unique-paths-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func uniquePathsWithObstacles(obstacleGrid [][]int) int {}
*/

// time complexity: O(n^2)
// space complexity: O(n)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	var m = len(obstacleGrid)
	var n = len(obstacleGrid[0])
	var d = make([]int, n)
	if obstacleGrid[0][0] == 1 {
		d[0] = 0
	} else {
		d[0] = 1
	}

	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			d[0] = 0
		}
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				d[j] = 0
			} else {
				d[j] = d[j] + d[j-1]
			}
		}
	}

	return d[n-1]
}
