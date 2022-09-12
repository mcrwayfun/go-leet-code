package main

import "fmt"

func findLength(str string, k int) int {
	if len(str) == 0 {
		return 0
	}

	var windowStart int
	var charFrequencyMap = make(map[byte]int, len(str))
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
		maxLength = max(maxLength, windowEnd-windowStart+1)
	}

	return maxLength
}

func main() {
	str := "araaci"
	k := 2

	str2 := "cbbebi"
	k2 := 3

	str3 := "araaci"
	k3 := 1

	fmt.Println(findLength(str, k))
	fmt.Println(findLength(str2, k2))
	fmt.Println(findLength(str3, k3))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
