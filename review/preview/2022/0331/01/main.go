package main

import "fmt"

func findLength(str string) int {
	if len(str) == 0 {
		return 0
	}

	var charFrequencyMap = make(map[byte]int, len(str))
	var maxLength int
	var windowStart int

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		if _, ok := charFrequencyMap[str[windowEnd]]; ok {
			windowStart = charFrequencyMap[str[windowEnd]] + 1
		}

		charFrequencyMap[str[windowEnd]] = windowEnd
		maxLength = max(maxLength, windowEnd-windowStart+1)
	}

	return maxLength
}

func main() {
	str := "aabccbb"
	fmt.Println(findLength(str))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
