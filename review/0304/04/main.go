package main

import "fmt"

func findLength(str string, k int) int {
	if len(str) == 0 {
		return 0
	}

	var windowStart int
	var maxLength int
	var charFrequencyMap = make(map[byte]int, len(str))

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
		maxLength = max(maxLength, windowEnd-windowStart+1)
	}

	return maxLength
}

func main() {
	// input: str=araaci, k=2
	// output: 4

	// inout: str=araaci, k=1
	// output: 2
	k := 2
	str := "araaci"
	characters := findLength(str, k)
	fmt.Println(characters)

	k = 1
	characters = findLength(str, k)
	fmt.Println(characters)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
