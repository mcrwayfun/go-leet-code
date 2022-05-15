package main

import "strconv"

/**
题目：
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

例如，121 是回文，而 123 不是。

示例 1：

输入：x = 121
输出：true
示例2：

输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3：

输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/palindrome-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func isPalindrome(x int) bool
*/

/**
解题思路1：将数字转换为字符，使用双指针从头和尾开始判断是否为回文

time complexity: O(m)
space complexity: O(1)
*/
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	left := 0
	right := len(str) - 1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

/**
解题思路2：将数字逆转，如果逆转后的数字和原数字相等，即是回文串。
数字 y=y*10+num，假设有数字12321，则有
1：num=tmp%10=1, y=0*10+num=1，tmp=tmp/10=1232
2：num=tmp%10=2, y=1*10+2=12, tmp=tmp/10=123
3：num=tmp%10=3, y=12*10+3=123, tmp=tmp/10=12
4：num=tmp%10=2，y=123*10+2=1232, tmp=tmp/10=1
5：num=tmp%10=1，y=1232*10+1=12321，tmp=tmp/10=0

time complexity: O(m)
space complexity: O(1)
*/
func isPalindromeV2(x int) bool {
	if x < 0 {
		return false
	}
	var tmp = x
	var y int
	for tmp != 0 {
		num := tmp % 10
		y = y*10 + num
		tmp = tmp / 10
	}
	return x == y
}
