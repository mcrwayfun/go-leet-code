package main

import "fmt"

/**
题目：Given a string, find the length of the longest substring
in it with no more than K distinct characters.

解析：给定一个字符串，找到不重复字符的个数不超过k的最长子串

举例：
Example 1:

Input: String="araaci", K=2
Output: 4
Explanation: The longest substring with no more than '2' distinct characters is "araa".
Example 2:

Input: String="araaci", K=1
Output: 2
Explanation: The longest substring with no more than '1' distinct characters is "aa".
Example 3:

Input: String="cbbebi", K=3
Output: 5
Explanation: The longest substrings with no more than '3' distinct characters are "cbbeb" & "bbebi".

模板：
func findLength(str string,k int) int
*/

/**
思路1：滑动窗口法
1：维护变量windowStart记录左侧窗口的初始值，使用map来记录字符出现的次数，map key要使用 byte。
使用maxLength来记录不重复的最大长度。
2：遍历字符串str
	2-1：使用map记录出现的字符的次数
	2-2：当map.size > k, 表示此时窗口中重复的字符已经超过了k，需要滑动窗口
		2-2-1：str[windowStart] 出现的次数减1
		2-2-2：str[windowStart] 出现次数为0时，将其从map中移除
	2-3：统计连续子串中不重复的最大长度 maxLength

步骤2-3为什么不在滑动窗口的时候统计？
需要看清楚 k 的条件，如果是 k == condition，因为是符合滑动窗口的条件，此时可以先将符合条件的
窗口记录，再滑动窗口。

像这道题目就是不一样的情况，因为滑动窗口后的长度才是我们需要统计的条件。
*/

func longestSubstringWithKDistinctCharacters(k int, str string) int {
	if len(str) == 0 {
		return 0
	}

	var windowStart int
	var duplicatedMap = make(map[byte]int, 0)
	var maxLength int

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		duplicatedMap[str[windowEnd]]++
		for len(duplicatedMap) > k {
			if _, ok := duplicatedMap[str[windowStart]]; ok {
				duplicatedMap[str[windowStart]]--
				if duplicatedMap[str[windowStart]] == 0 {
					delete(duplicatedMap, str[windowStart])
				}
			}
			windowStart++
		}
		maxLength = max(maxLength, windowEnd-windowStart+1)
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// input: str=araaci, k=2
	// output: 4

	// inout: str=araaci, k=1
	// output: 2
	k := 2
	str := "araaci"
	characters := longestSubstringWithKDistinctCharacters(k, str)
	fmt.Println(characters)

	k = 1
	characters = longestSubstringWithKDistinctCharacters(k, str)
	fmt.Println(characters)
}
