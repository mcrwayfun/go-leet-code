package main

import "fmt"

// 给定一个字符串，找到不重复字符的个数不超过k的最长子串
func findLength(str string, k int) int {
	if len(str) == 0 {
		return 0
	}

	var charFrequencyMap = make(map[byte]int, len(str))
	var windowStart int
	var maxLength int

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		charFrequencyMap[str[windowEnd]]++

		for len(charFrequencyMap) > k {
			if _, ok := charFrequencyMap[str[windowStart]]; ok {
				charFrequencyMap[str[windowStart]]--
				if charFrequencyMap[str[windowStart]] == 0 {
					delete(charFrequencyMap, str[windowStart])
				}
			}
			windowStart++
		}

		if windowEnd-windowStart+1 > maxLength {
			maxLength = windowEnd - windowStart + 1
		}
	}

	return maxLength
}

func main() {
	str := "araaci"
	k := 2
	fmt.Println(findLength(str, k))
}
