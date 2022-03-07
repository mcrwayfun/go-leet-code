package main

import "fmt"

/**
题目：Given a string and a pattern, find all anagrams of the pattern in the given string.
Write a function to return a list of starting indices
of the anagrams of the pattern in the given string.

解析：给一个字符串和pattern，找到所有符合pattern的排列组合。编写一个方法返回
给定字符串中符合pattern的排列组合的"起始索引列表"。

举例：
Example 1: pq和qp都符合要求，所以这两个字符串的起始下标分别是1,2

Input: String="ppqp", Pattern="pq"
Output: [1, 2]
Explanation: The two anagrams of the pattern in the given string are "pq" and "qp".

Example 2:

Input: String="abbcabc", Pattern="abc"
Output: [2, 3, 4]
Explanation: The three anagrams of the pattern in the given string are "bca", "cab", and "abc".

func findStringAnagrams(str string, pattern string) []int
*/

/**
解题思路1：使用滑动窗口法
1：维护一个map，key记录的是字符，value记录的是字符出现的次数。使用match来记录有字符匹配上了pattern。
使用一个slice来记录完全匹配patter的子串的first index
2：遍历patter，map.put(key)++
3：遍历str
	3-1：if map.contains(key), then--
	3-2：if map.key.value == 0, then matched++
	3-3：if map.size == matched, 意味着完全匹配, 记录first index=windowStart
	3-4：if windowEnd-windowStart+1 >= len(pattern) mean we should shirk window
		3-4-1：if map.key.value == 0, matched--
		3-4-2：map.contains(key), then++
		3-4-3：windowStart++

time complexity: O(n)
space complexity: O(k), k is the length of pattern
*/
func findStringAnagrams(str string, pattern string) []int {
	if len(str) == 0 {
		return []int{}
	}

	var charFrequencyMap = make(map[byte]int, len(pattern))
	for i := 0; i < len(pattern); i++ {
		charFrequencyMap[pattern[i]]++
	}

	var windowStart int
	var matched int
	var retIndices []int

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]
		if _, ok := charFrequencyMap[rightChar]; ok {
			charFrequencyMap[rightChar]--
			if charFrequencyMap[rightChar] == 0 {
				matched++
			}
		}

		if len(charFrequencyMap) == matched {
			retIndices = append(retIndices, windowStart)
		}

		if windowEnd-windowStart+1 >= len(pattern) {
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

	return retIndices
}

func main() {
	str := "ppqp"
	pattern := "pq"

	str2 := "abbcabc"
	pattern2 := "abc"

	fmt.Println(findStringAnagrams(str, pattern))
	fmt.Println(findStringAnagrams(str2, pattern2))
}
