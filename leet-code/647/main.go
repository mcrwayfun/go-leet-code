package main

import "fmt"

/**
题目：
给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。

回文字符串 是正着读和倒过来读一样的字符串。

子字符串 是字符串中的由连续字符组成的一个序列。

具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

示例 1：

输入：s = "abc"
输出：3
解释：三个回文子串: "a", "b", "c"
示例 2：

输入：s = "aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/palindromic-substrings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：
func countSubstrings(s string) int
*/

/**
解题思路：假设现在有字符aba，它的回文子串有a、b、a、aba，
按照如此可以发现的规律有，当存在对比两个字符i和j是否回文串，s(i,j)
1：i==j，则只有一个字符，是回文串
2：i+1==j，相邻的两个字符相等，是回文串
3：s(i)==s(j) && s(i+1)==s(j-1)
这道题是利用动态规划的思路，i从n-1->0, j从i->n-1，从最小解扩散到全部解。
申明一个二维数组来存储已经访问过的路径是否为回文串。

time complexity: O(n^2)
space complexity: O(n^2)
*/
func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}

	var n = len(s)
	var count int

	visited := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		visited[i] = make([]bool, len(s))
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j <= n-1; j++ {
			if i == j {
				visited[i][j] = true
			} else if i+1 == j {
				if s[i] == s[j] {
					visited[i][j] = true
				}
			} else {
				if s[i] == s[j] && visited[i+1][j-1] {
					visited[i][j] = true
				}
			}

			if visited[i][j] == true {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(countSubstrings("aba"))
}
