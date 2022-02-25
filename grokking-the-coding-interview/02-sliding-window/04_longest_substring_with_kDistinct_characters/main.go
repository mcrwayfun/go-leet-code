package main

import "fmt"

func longestSubstringWithKDistinctCharacters(k int, str string) int {
	if len(str) == 0 {
		return 0
	}

	var windowStart int
	var duplicatedMap = make(map[byte]int, 0)
	var maxLength int

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		duplicatedMap[str[windowEnd]]++
		if len(duplicatedMap) > k {
			maxLength = max(maxLength, windowEnd-windowStart+1)
			if _, ok := duplicatedMap[str[windowStart]]; ok {
				duplicatedMap[str[windowStart]]--
			}
			windowStart++
		}
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
	k := 2
	str := "araaci"
	characters := longestSubstringWithKDistinctCharacters(k, str)
	fmt.Println(characters)

	k = 1
	characters = longestSubstringWithKDistinctCharacters(k, str)
	fmt.Println(characters)
}