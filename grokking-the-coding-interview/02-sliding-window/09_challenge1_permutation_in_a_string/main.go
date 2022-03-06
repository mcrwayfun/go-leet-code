package main

import "fmt"

/**
题目：Given a string and a pattern, find out if the string contains any permutation of the pattern.

解析：给定一个字符串和pattern，判断字符串中是否包含这个pattern的任何排列组合。

举例：
Example 1:

Input: String="oidbcaf", Pattern="abc"
Output: true
Explanation: The string contains "bca" which is a permutation of the given pattern.
Example 2:

Input: String="odicf", Pattern="dc"
Output: false
Explanation: No permutation of the pattern is present in the given string as a substring.
Example 3:

Input: String="bcdxabcdy", Pattern="bcdyabcdx"
Output: true
Explanation: Both the string and the pattern are a permutation of each other.
Example 4:

Input: String="aaacb", Pattern="abc"
Output: true
Explanation: The string contains "acb" which is a permutation of the given pattern.

模板：
func findPermutation(str string, pattern string) bool
*/

/**
解法1：暴力破解法
1：使用map来存储pattern中字符出现的次数
2：遍历字符str，判断是否存在符合条件的子串
	2-1：char 在map中出现则将key对应的value减一
	2-2：如果 key对应的value为0，则执行delete
	2-3：如果map.size==0, 则表示存在子串符合排列组合，否则不符合

time complexity: O(n^2)
space complexity: O(n)
*/
func findPermutation(str string, pattern string) bool {
	if len(str) == 0 {
		return false
	}

	var charFrequencyMap = make(map[byte]int, len(pattern))
	for i := 0; i < len(pattern); i++ {
		charFrequencyMap[pattern[i]]++
	}

	for i := 0; i < len(str); i++ {
		var tmpCharFrequencyMap = make(map[byte]int, len(pattern)) // space: O(n)
		copyMap(tmpCharFrequencyMap, charFrequencyMap)             // time: O(n)

		for j := i; j < len(str); j++ { // time: O(n)
			char := str[j]
			if _, ok := charFrequencyMap[char]; ok {
				tmpCharFrequencyMap[char]--
				if tmpCharFrequencyMap[char] == 0 {
					delete(tmpCharFrequencyMap, char)
				}

				if len(tmpCharFrequencyMap) == 0 {
					return true
				}
			} else {
				break
			}
		}
	}

	return false
}

/**
解法2：滑动窗口法
1：使用map来统计pattern中字符出现的次数，使用全局变量matched来表示满足pattern的字符个数
2：遍历字符串str
	2-1：如果当前字符char存在于map中，则key对应的value--
	2-2：如果key对应的value为0，则表示有个字符匹配，matched++
3：如果 matched == map.size, 则表示存在完全匹配的字符，可以返回true
4：滑动窗口的时机：windowEnd-windowStart+1>=len(pattern)，注意这里是>=，因为窗口长度和pattern长度完全一致的情况下，
还没有找到排列组合，也需要滑动窗口。
	4-1：if map[str[windowStart]] == 0, matched--
	4-2：map[str[windowStart]]++
*/
func findPermutation2(str string, pattern string) bool {
	if len(str) == 0 {
		return false
	}

	var charFrequencyMap = make(map[byte]int, len(pattern))
	for i := 0; i < len(pattern); i++ {
		charFrequencyMap[pattern[i]]++
	}

	var windowStart int
	var matched int

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]
		if _, ok := charFrequencyMap[rightChar]; ok {
			charFrequencyMap[rightChar]--
			if charFrequencyMap[rightChar] == 0 {
				matched++
			}
		}

		if matched == len(charFrequencyMap) {
			return true
		}

		if windowEnd-windowStart+1 >= len(pattern) {
			leftChar := str[windowStart]
			if _, ok := charFrequencyMap[leftChar]; ok {
				if charFrequencyMap[leftChar] == 0 {
					matched--
				}
				charFrequencyMap[leftChar]++ // 返回移除的参数到 charFrequencyMap中
			}
			windowStart++
		}
	}

	return false
}

func main() {
	str := "oidbcaf"
	pattern := "abc"

	str2 := "odicf"
	pattern2 := "dc"

	str3 := "bcdxabcdy"
	pattern3 := "bcdyabcdx"

	str4 := "aaacb"
	pattern4 := "abc"

	fmt.Println(findPermutation(str, pattern))
	fmt.Println(findPermutation(str2, pattern2))
	fmt.Println(findPermutation(str3, pattern3))
	fmt.Println(findPermutation(str4, pattern4))

	fmt.Println(findPermutation2(str, pattern))
	fmt.Println(findPermutation2(str2, pattern2))
	fmt.Println(findPermutation2(str3, pattern3))
	fmt.Println(findPermutation2(str4, pattern4))
}

func copyMap(desc, src map[byte]int) {
	for k, v := range src {
		desc[k] = v
	}
}
