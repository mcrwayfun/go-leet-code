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
			leftChar := str[windowStart]
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
	str := "araaci"
	k := 2

	str2 := "araaci"
	k2 := 1

	str3 := "cbbebi"
	k3 := 3

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
