package main

/**
丑数 就是只包含质因数2、3 和 5的正整数。

给你一个整数 n ，请你判断 n 是否为 丑数 。如果是，返回 true ；否则，返回 false 。

示例 1：
输入：n = 6
输出：true
解释：6 = 2 × 3

示例 2：
输入：n = 1
输出：true
解释：1 没有质因数，因此它的全部质因数是 {2, 3, 5} 的空集。习惯上将其视作第一个丑数。

示例 3：
输入：n = 14
输出：false
解释：14 不是丑数，因为它包含了另外一个质因数7 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/ugly-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func isUgly(n int) bool {}
*/

/**
解题思路：一个丑数能够不断的被2/3/5整除，如果最后是1则认为是丑数，否则不是

time complexity: O(m+n+l)
space complexity: O(1)
*/
func isUgly(n int) bool {
	var res = n
	if res <= 0 {
		return false
	}

	for res%2 == 0 {
		res = res / 2
	}
	for res%3 == 0 {
		res = res / 3
	}
	for res%5 == 0 {
		res = res / 5
	}
	return res == 1
}
