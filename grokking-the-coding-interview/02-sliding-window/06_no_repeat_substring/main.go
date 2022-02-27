package main

import "fmt"

/**
题目：Given a string, find the length of the longest substring which has no repeating characters.

解析：给定一个字符串，找到不包含任何重复字符的最长子串。

举例：
Example 1:

Input: String="aabccbb"
Output: 3
Explanation: The longest substring without any repeating characters is "abc".
Example 2:

Input: String="abbbb"
Output: 2
Explanation: The longest substring without any repeating characters is "ab".
Example 3:

Input: String="abccde"
Output: 3
Explanation: Longest substrings without any repeating characters are "abc" & "cde".

模板：
func findLength(str string) int
*/

/**
思路1：使用暴力破解法
1：维护一个全量变量maxLength记录最大长度，维护一个局部变量map，用来记录字符出现的次数
2：外层循环遍历str，内层循环记录字符出现的次数，若有重复则计算一次maxLength

time complexity: O(n^k) k is the no repeat substring length
space complexity: O(k)
*/
func findLength(str string) int {
	if len(str) == 0 {
		return 0
	}

	var maxLength int
	for i := 0; i < len(str); i++ {
		var charFrequencyMap = make(map[byte]int, 0)
		for j := i; j < len(str); j++ {
			if _, ok := charFrequencyMap[str[j]]; ok {
				// 以下两种方式均可以
				// maxLength = max(maxLength, j-i+1-1)
				// maxLength = max(maxLength, len(charFrequencyMap))
				maxLength = max(maxLength, j-i)
				break
			}
			charFrequencyMap[str[j]]++
		}
	}

	return maxLength
}

/**
思路2：滑动窗口法
1：使用全局变量maxLength来记录符合条件的子串的最大长度，使用map来记录字符出现的下标。
使用windowStart来记录窗口的起始位置
2：遍历字符串str，
	2-1：如果字符在map中存在的话，需要滑动窗口。此时将 windowStart = 该字符上次出现的位置 + 1
	2-2：如果字符在map中没有出现过的话，则 map[key] = windowEnd
	2-3：maxLength = windowEnd - windowStart + 1（此时是符合条件的，子串中没有重复的）

time complexity: O(n)
space complexity: O(1)
*/
func findLength2(str string) int {
	if len(str) == 0 {
		return 0
	}

	var windowStart int
	var charIndexMap = make(map[byte]int, len(str))
	var maxLength int

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		char := str[windowEnd]
		if _, ok := charIndexMap[char]; ok {
			windowStart = charIndexMap[char] + 1 // need to shrink window
		}
		charIndexMap[char] = windowEnd
		maxLength = max(maxLength, windowEnd-windowStart+1)
	}

	return maxLength
}

func main() {
	str := "aabccbb"
	str2 := "abbbb"
	str3 := "abccde"
	fmt.Println(findLength(str))
	fmt.Println(findLength(str2))
	fmt.Println(findLength(str3))

	fmt.Println(findLength2(str))
	fmt.Println(findLength2(str2))
	fmt.Println(findLength2(str3))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
