package main

/**
两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。

给你两个整数 x 和 y，计算并返回它们之间的汉明距离。

示例 1：
输入：x = 1, y = 4

输出：2
解释：
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑
上面的箭头指出了对应二进制位不同的位置。
示例 2：

输入：x = 3, y = 1
输出：1

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/hamming-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func hammingDistance(x int, y int) int {}
*/

// time complexity: O(m)
// space complexity: O(1)
func hammingDistance(x int, y int) int {
	return numberOfOne(x ^ y)
}

func numberOfOne(n int) int {
	var count int
	for n != 0 {
		count++
		n &= n - 1
	}
	return count
}
