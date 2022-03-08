package main

import "fmt"

/**
题目：Given a string and a pattern, find the smallest substring in the given string
which has all the characters of the given pattern.

解析：给出一个字符串和一个模式，在给定的字符串中找到最小的子串。
该子串包含了给定模式中的所有字符。

举例：
Example 1:

Input: String="aabdec", Pattern="abc"
Output: "abdec"
Explanation: The smallest substring having all characters of the pattern is "abdec"
Example 2:

Input: String="abdabca", Pattern="abc"
Output: "abc"
Explanation: The smallest substring having all characters of the pattern is "abc".
Example 3:

Input: String="adcad", Pattern="abc"
Output: ""
Explanation: No substring in the given string has all characters of the pattern.

模板：
func findSubstring(str string, pattern string) string
*/

/**
解决思路1：滑动窗口法。
先获取到符合pattern要求的字符，然后再通过移除掉重复的字符来求得 minLength。'

1：使用map来统计pattern中字符出现的次数，使用matched来记录匹配到pattern中字符的次数。
2：使用minLength=len(arr)+1来记录符合条件的最短长度，由于要返回的是子串，所以还需要有变量subStart来记录子串的起始位置。
3：for pattern, map.put(char)++
4：for str,
	4-1：if map.contains(rightChar) then map[rightChar]--
	4-2：if map[rightChar] == 0 then matched++ （如果pattern存在重复的字符，这里要使用>=，而不是==）
	4-3: for matched == len(pattern) then we should shrink window and remove the duplicated characters
		4-3-1：if minLength > windowSize we should re-calculate the minLength and subStart:
	minLength = windowSize and subStart = windowStart
		4-3-2：if map.contains(leftChar)
			4-3-2-1: if map[leftChar] == 0 then matched--
			4-3-2-2: map[leftChar]++
		4-3-3: windowStart++
*/
func findSubstring(str string, pattern string) string {
	if len(str) == 0 {
		return ""
	}

	var charFrequencyMap = make(map[byte]int, len(pattern))
	for i := 0; i < len(pattern); i++ {
		charFrequencyMap[pattern[i]]++
	}

	var windowStart int
	var minLength = len(str) + 1
	var subStart int
	var matched int

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]

		if _, ok := charFrequencyMap[rightChar]; ok {
			charFrequencyMap[rightChar]--
			if charFrequencyMap[rightChar] >= 0 { // 如果pattern存在重复的字符，这里要使用>=，而不是==
				matched++
			}
		}

		for matched == len(pattern) {
			// we should shrink the window and remove the duplicated character
			if minLength > windowEnd-windowStart+1 {
				minLength = windowEnd - windowStart + 1
				subStart = windowStart
			}

			leftChar := str[windowStart]
			if _, ok := charFrequencyMap[leftChar]; ok {
				if charFrequencyMap[leftChar] == 0 {
					matched--
				}
				charFrequencyMap[leftChar]++
			}

			windowStart++
		}
	}

	if minLength == len(str)+1 || (subStart+minLength) > len(str) {
		return ""
	}

	return str[subStart : subStart+minLength]
}

func main() {
	str := "aabdec"
	pattern := "abc"
	fmt.Println(findSubstring(str, pattern))

	str = "abdabca"
	pattern = "abc"
	fmt.Println(findSubstring(str, pattern))

	str = "adcad"
	pattern = "abc"
	fmt.Println(findSubstring(str, pattern))

	str = "aabdec"
	pattern = "aabc"
	fmt.Println(findSubstring(str, pattern))
}
