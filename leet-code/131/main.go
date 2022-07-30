package main

/**
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

回文串 是正着读和反着读都一样的字符串。
示例 1：
输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]

示例 2：
输入：s = "a"
输出：[["a"]]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/palindrome-partitioning
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func partition(s string) [][]string {}
*/

/**
1：利用动态规划来判断字符串s的回文状态, 表现为二维数组d
2：利用回溯的思想来进行排列组合，求出所有组合的回文子串
申明两个指针start=0，end=len(s)-1。指针start遍历到end，
	2-1：假设s[start]是个回文串，则可以直接从 start+1 -> end
	2-2：start+1 -> end 本身和步骤2-1一样，是一个回溯的过程
 */
// time complexity: O(2^n)
// space complexity: O(n^2)
func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}

	var n = len(s)
	var d = make([][]bool, n)
	for i := 0; i < n; i++ {
		d[i] = make([]bool, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				d[i][j] = true
			} else if i+1 == j {
				d[i][j] = s[i] == s[j]
			} else {
				d[i][j] = s[i] == s[j] && d[i+1][j-1]
			}
		}
	}

	var ret [][]string
	var ele []string
	helper(s, 0, d, &ret, &ele)
	return ret
}

func helper(s string, start int, d [][]bool,
	ret *[][]string, ele *[]string) {
	if start >= len(s) {
		*ret = append(*ret, append([]string{}, *ele...))
	} else {
		// 回溯
		for end := start; end < len(s); end++ {
			if d[start][end] { // 当前是回文串
				*ele = append(*ele, s[start:end+1])
				helper(s, end+1, d, ret, ele)
				*ele = (*ele)[:len(*ele)-1]
			}
		}
	}
}
