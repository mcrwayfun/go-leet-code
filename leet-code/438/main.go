package main

/**
给定两个字符串s和 p，找到s中所有p的异位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

示例1:
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。

示例 2:
输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/find-all-anagrams-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// time complexity: O(n)
// space complexity: O(n)
func findAnagrams(s string, p string) []int {
	if len(s) == 0 || len(p) == 0 {
		return []int{}
	}

	var charFrequencyMap = make(map[byte]int, len(p))
	for i := 0; i < len(p); i++ {
		charFrequencyMap[p[i]]++
	}

	var start int
	var res []int
	var matched int
	for end := 0; end < len(s); end++ {
		if _, ok := charFrequencyMap[s[end]]; ok {
			charFrequencyMap[s[end]]--
			if charFrequencyMap[s[end]] == 0 {
				matched++
			}
		}

		// why use matched == len(charFrequencyMap) but not to use matched == len(p)
		// because while the p = aabc, the matched would be 3 not to 4
		if matched == len(charFrequencyMap) {
			res = append(res, start)
		}

		// who not use >= but not to >
		// because will cal start multiple times
		if end-start+1 >= len(p) {// should shrink the window
			if _, ok := charFrequencyMap[s[start]]; ok {
				if charFrequencyMap[s[start]] == 0 {
					matched--
				}
				charFrequencyMap[s[start]]++
			}
			start++
		}
	}

	return res
}
