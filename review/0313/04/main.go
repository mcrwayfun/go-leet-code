package main

import "fmt"

func findLength(str string, k int) int {
	if len(str) == 0 {
		return 0
	}

	var windowStart int
	var charFrequencyMap = make(map[byte]int, len(str))
	var maxRepeatCount int
	var maxLength int

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]
		charFrequencyMap[rightChar]++
		maxRepeatCount = max(maxRepeatCount, charFrequencyMap[rightChar])

		if windowEnd-windowStart+1-maxRepeatCount > k {
			leftChar := str[windowStart]
			charFrequencyMap[leftChar]--
			windowStart++
		}

		maxLength = max(maxLength, windowEnd-windowStart+1)
	}

	return maxLength
}

func main() {
	str := "aabccbb"
	k := 2

	str2 := "abbcb"
	k2 := 1

	str3 := "abccde"
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
