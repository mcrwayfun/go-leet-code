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
在 Go 语言中，可以使用切片来生成「杨辉三角」。

以下是代码实现：

func generate(numRows int) [][]int {
    triangle := make([][]int, numRows)
    for i := 0; i < numRows; i++ {
        triangle[i] = make([]int, i+1)
        triangle[i][0] = 1
        triangle[i][i] = 1
        for j := 1; j < i; j++ {
            triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
        }
    }
    return triangle
}
在这个代码中，我们首先声明一个切片 triangle，并在循环中创建二维切片，其中，第一维代表行数，第二维代表列数。

在每一行中，我们首先将第一个元素设置为 1，最后一个元素也设置为 1。其他元素的值是上一行左上方元素与右上方元素之和。

例如：

numRows = 5

1
1 1
1 2 1
1 3 3 1
1 4 6 4 1
这样就可以生成「杨辉三角」了。

time complexity: O(n^2)
space complexity: O(1)
*/
func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	return triangle
}

func main() {
	fmt.Println(generate(5))
}
