package main

import "fmt"

// 给定一个字符串，找到不重复字符的个数不超过k的最长子串
// Input: String="araaci", K=2
// Output: 4
func findLength(str string, k int) int {
	if len(str) == 0 {
		return 0
	}

	var charFrequencyMap = make(map[byte]int)
	var windowStart int
	var maxLength int

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		charFrequencyMap[str[windowEnd]]++

		for len(charFrequencyMap) > k {
			charFrequencyMap[str[windowStart]]--
			if charFrequencyMap[str[windowStart]] == 0 {
				delete(charFrequencyMap, str[windowStart])
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

	str = "araaci"
	k = 1
	fmt.Println(findLength(str, k))

	str = "cbbebi"
	k = 3
	fmt.Println(findLength(str, k))
}
