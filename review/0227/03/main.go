package main

import "fmt"

// time complexity: O(n)
// space complexity: O(k)
func findLength(str string, k int) int {
	if len(str) == 0 {
		return 0
	}

	var windowStart int
	var charFrequencyMap = make(map[byte]int, len(str))
	var maxLength int // return 0 if no max length

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
		// 进入for循环的是不符合要求的，因为元素个数大于k
		// 所以应该在for循环外求 maxLength
		maxLength = max(maxLength, windowEnd-windowStart+1)
	}

	return maxLength
}

func main() {
	str := "araaci"
	k := 2

	fmt.Println(findLength(str, k))
	fmt.Println(findLength(str, 1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
