package main

import "fmt"

func findLength(strs []string) int {
	if len(strs) == 0 {
		return 0
	}

	var windowStart int
	var maxLength int
	var charFrequencyMap = make(map[string]int, len(strs))

	k := 2

	for windowEnd := 0; windowEnd < len(strs); windowEnd++ {
		charFrequencyMap[strs[windowEnd]]++
		for len(charFrequencyMap) > k {
			if _, ok := charFrequencyMap[strs[windowStart]]; ok {
				charFrequencyMap[strs[windowStart]]--
				if charFrequencyMap[strs[windowStart]] == 0 {
					delete(charFrequencyMap, strs[windowStart])
				}
			}
			windowStart++
		}
		maxLength = max(maxLength, windowEnd-windowStart+1)
	}

	return maxLength
}

func main() {
	strs := []string{"A", "B", "C", "A", "C"}
	strs2 := []string{"A", "B", "C", "B", "B", "C"}
	fmt.Println(findLength(strs))
	fmt.Println(findLength(strs2))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
