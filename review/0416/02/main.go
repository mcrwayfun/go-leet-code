package main

import "fmt"

func findPermutation(str string, pattern string) bool {
	if len(str) == 0 || len(pattern) == 0 {
		return false
	}

	var charFrequencyMap = make(map[byte]int, len(pattern))
	for i := 0; i < len(pattern); i++ {
		charFrequencyMap[pattern[i]]++
	}

	var windowStart int
	var matched int // 记录匹配的次数

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		curChar := str[windowEnd]
		if _, ok := charFrequencyMap[curChar]; ok {
			charFrequencyMap[curChar]--
			// pattern存在重复字符的时候，也算是匹配
			if charFrequencyMap[curChar] >= 0 {
				matched++
			}
		}

		if matched == len(pattern) {
			return true
		}

		// we should shrink the window
		if windowEnd-windowStart+1 >= len(pattern) {
			startChar := str[windowStart]
			if _, ok := charFrequencyMap[startChar]; ok {
				if charFrequencyMap[startChar] >= 0 {
					matched--
				}
				charFrequencyMap[startChar]++
			}
			windowStart++
		}
	}

	return false
}

func main() {
	str := "oidbcaf"
	pattern := "abc"

	fmt.Println(findPermutation(str, pattern))

	str = "bcdxabcdy"
	pattern = "bcdyabcdx"
	fmt.Println(findPermutation(str, pattern))

	str = "odicf"
	pattern = "dc"
	fmt.Println(findPermutation(str, pattern))
}
