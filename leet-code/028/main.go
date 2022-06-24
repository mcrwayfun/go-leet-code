package main

/**
实现strStr()函数。

给你两个字符串haystack 和 needle ，
请你在 haystack 字符串中找出 needle 字符串
出现的第一个位置（下标从 0 开始）。如果不存在，则返回 -1 。

说明：
当needle是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
对于本题而言，当needle是空字符串时我们应当返回 0 。
这与 C 语言的strstr()以及 Java 的indexOf()定义相符。

示例 1：
输入：haystack = "hello", needle = "ll"
输出：2

示例 2：
输入：haystack = "aaaaa", needle = "bba"
输出：-1

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/implement-strstr
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。1

模板: func strStr(haystack string, needle string) int {}
*/

/**
time complexity: O(m+n) * m
space complexity: O(1)
*/
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	var n = len(haystack)
	var m = len(needle)
	for i := 0; i <= n-m; i++ {
		var k = i
		var j = 0
		for ; j < m && haystack[k] == needle[j]; {
			j++
			k++
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}
