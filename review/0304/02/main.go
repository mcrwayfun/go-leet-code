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
		charFrequencyMap[str[windowEnd]] = windowEnd
		if len(charFrequencyMap) > k {
			windowStart = charFrequencyMap[str[windowEnd]] + 1
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

	fmt.Println(findLength(str, k))
	fmt.Println(findLength(str2, k2))

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
