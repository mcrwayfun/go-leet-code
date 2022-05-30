package main

/**
给你一个字符串 s，找到 s 中最长的回文子串。

示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

示例 2：
输入：s = "cbbd"
输出："bb"

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func longestPalindrome(s string) string {}
*/

/**
解法1：使用动态规划来，通过求小解来求得最终解
假设字符为abcbad, 最长回文串为abcba，对于s(i,j)
1：i==j, 则肯定是回文子串
2：i+1==j, s(i)==s(j)，则也是回文子串
3：else s(i)==s(j) && s(i+1,j-1)，则是回文子串

通过申明一个二维数组来存储已经遍历过的路径，
i: 从大到小：n-1 -> 0, 因为需要知道i+1的值
j: 从小到大：i -> n-1, 因为需要知道j-1的值

time complexity: O(n^2)
space complexity: O(n^2)
*/
func longestPalindromeDP(s string) string {
	if len(s) == 0 {
		return ""
	}

	// 初始化二维数组
	var visited = make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		visited[i] = make([]bool, len(s))
	}

	var n = len(s)
	var subStart int
	var maxLen int
	for i := n - 1; i >= 0; i-- {
		for j := i; j <= n-1; j++ {
			if i == j {
				visited[i][j] = true
			} else if i+1 == j && s[i] == s[j] {
				visited[i][j] = true
			} else {
				if s[i] == s[j] && visited[i+1][j-1] {
					visited[i][j] = true
				}
			}

			if visited[i][j] && (j-i+1) > maxLen {
				subStart = i
				maxLen = j - i + 1
			}
		}
	}

	return s[subStart : subStart+maxLen]
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	var subStart int
	var maxLen int
	for i := 0; i < len(s); i++ {
		length1 := expand(s, i, i)
		length2 := expand(s, i, i+1)
		length := max(length1, length2)

		if length > maxLen {
			subStart = i - (length-1)/2
			maxLen = length
		}
	}
	return s[subStart : subStart+maxLen]
}

func expand(s string, left, right int) int {
	var count int
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
		count++
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
