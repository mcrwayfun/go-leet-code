package main

import "fmt"

/**
题目：Given a string with lowercase letters only, if you are allowed to replace no more than ‘k’
letters with any letter, find the length of the longest substring
having the same letters after replacement.

解析：给定一个全小写的字符串，可以替换其中n个字符（n<=k), 找到替换后，全部都是相同字符的最长子串。

举例：
Example 1:

Input: String="aabccbb", k=2
Output: 5
Explanation: Replace the two 'c' with 'b' to have a longest repeating substring "bbbbb".
Example 2:

Input: String="abbcb", k=1
Output: 4
Explanation: Replace the 'c' with 'b' to have a longest repeating substring "bbbb".
Example 3:

Input: String="abccde", k=1
Output: 3
Explanation: Replace the 'b' or 'd' with 'c' to have the longest repeating substring "ccc".

题目模板：
func findLength(str string, k int) int
*/

/**
解题思路：使用滑动窗口的思想，这里的关键是 滑动窗口的时机
1：使用map来统计元素出现的次数，维护一个 maxRepeatCount 来统计当前窗口中同个元素出现的最大个数。
2：滑动窗口的时机：windowEnd-windowStart+1-maxRepeatCount > k，举个例子：对于aabccbb，k=2
假设当前窗口为 aabcc, 此时：
	windowEnd=4, windowStart=0, windowSize=5, maxRepeatCount=2(aa/cc), k=2
则此时已经到达滑动窗口的时机，因为此时将窗口中最大的重复值为maxRepeatCount, 替换k个字符后，就是最大的相同字符子串了。

time complexity: O(n)
space complexity: O(n)
*/
func findLength(str string, k int) int {
	if len(str) == 0 {
		return 0
	}

	var maxLength int
	var maxRepeatCount int
	var windowStart int
	var charFrequencyMap = make(map[byte]int, len(str))

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		charFrequencyMap[str[windowEnd]]++
		maxRepeatCount = max(maxRepeatCount, charFrequencyMap[str[windowEnd]])

		if windowEnd-windowStart+1-maxRepeatCount > k {
			if _, ok := charFrequencyMap[str[windowStart]]; ok {
				charFrequencyMap[str[windowStart]]--
			}
			windowStart++
		}

		maxLength = max(maxLength, windowEnd-windowStart+1)
	}

	return maxLength
}

func main() {
	str := "aabccbb"
	k := 2

	fmt.Println(findLength(str, k))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
