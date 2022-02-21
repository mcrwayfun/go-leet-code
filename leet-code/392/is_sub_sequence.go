package _92

/**
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
你可以认为 s 和 t 中仅包含英文小写字母。字符串 t 可能会很长（长度 ~= 500,000），而 s 是个短字符串（长度 <=100）。
字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。
（例如，"ace"是"abcde"的一个子序列，而"aec"不是）

链接：https://leetcode-cn.com/problems/is-subsequence
time complexity:O(n)
space complexity:O(1)
*/
func isSubsequence(s string, t string) bool {
	// s和t均为空字符串,s是t的子序列
	if len(t) == 0 && len(s) == 0 {
		return true
	}
	// s为空字符串,而t不是
	if len(s) == 0 {
		return true
	}
	// t为空字符串,而s不是
	if len(t) == 0 {
		return false
	}

	m, n := 0, 0
	for m < len(t) && n < len(s){
		if t[m] == s[n] {
			m++
			n++
		} else {
			m++
		}
	}

	return n == len(s)
}
