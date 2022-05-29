package main

/**
给定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。

示例1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
    请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func lengthOfLongestSubstring(s string) int
*/

/**
使用滑动窗口的方法来解决该问题
1：申明一个map，用来存储字符出现的下标
2：维护一个滑动窗口，窗口中是当前不含有重复字符的最长子串。
	2-1：记录当前元素出现的下标
	2-2：如果当前元素已经在map中出现过，说明该移动窗口了，将start移动到当前
元素的下标的下一个位置
	2-3：滑动窗口结束时（或者每一次循环），因为当前窗口肯定没有重复的子元素，可以
统计下当前窗口的不含有重复字符的最长子串的长度。
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var noDuplicatedMap = make(map[byte]int, len(s))
	var start int
	var maxLength int

	for end := 0; end < len(s); end++ {
		if _, ok := noDuplicatedMap[s[end]]; ok {
			start = max(noDuplicatedMap[s[end]]+1, start)
		}
		noDuplicatedMap[s[end]] = end
		maxLength = max(end-start+1, maxLength)
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
