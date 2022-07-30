package main

/**
给定一个包含大写字母和小写字母的字符串s，返回通过这些字母构造成的 最长的回文串。
在构造过程中，请注意 区分大小写 。比如"Aa"不能当做一个回文字符串。

示例 1:
输入:s = "abccccdd"
输出:7
解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。

示例 2:
输入:s = "a"
输入:1

示例 3:
输入:s = "bb"
输入: 2

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-palindrome
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// time complexity: O(n)
// space complexity: O(m)
func longestPalindrome(s string) int {
	var charCountMap = make(map[byte]int, 256)
	for i := 0; i < len(s); i++ {
		charCountMap[s[i]]++
	}
	var oddNum int
	for _, count := range charCountMap {
		if count%2 != 0 {
			oddNum++
		}
	}

	// 奇数的字符组中，有一个是不能用的
	// 如果都是偶数组，则oddNum - 1 为-1，需要和0比较
	unused := max(oddNum-1, 0)
	return len(s) - unused
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
