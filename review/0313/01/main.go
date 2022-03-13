package main

import "fmt"

func findLength(strs []string) int {
	if len(strs) == 0 {
		return 0
	}

	var windowStart int
	var maxLength int
	var charFrequencyMap = make(map[string]int)

	for windowEnd := 0; windowEnd < len(strs); windowEnd++ {
		charFrequencyMap[strs[windowEnd]]++
		if len(charFrequencyMap) > 2 {
			leftChar := strs[windowStart]
			if _, ok := charFrequencyMap[leftChar]; ok {
				charFrequencyMap[leftChar]--
				if charFrequencyMap[leftChar] == 0 {
					delete(charFrequencyMap, leftChar)
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
	fmt.Println(findLength(strs))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
