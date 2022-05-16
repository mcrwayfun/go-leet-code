package main

import "fmt"

func findLength(str string, k int) int {
	if len(str) == 0 {
		return 0
	}

	var start int
	var maxRepeatCount int
	var charFrequencyMap = make(map[byte]int, len(str))
	var maxLength int

	for end := 0; end < len(str); end++ {
		charFrequencyMap[str[end]]++
		if charFrequencyMap[str[end]] > maxRepeatCount {
			maxRepeatCount = charFrequencyMap[str[end]]
		}

		for end-start+1-maxRepeatCount > k {
			charFrequencyMap[str[start]]--
			start++
		}

		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
	}

	return maxLength
}

func main() {
	str := "aabccbb"
	k := 2
	fmt.Println(findLength(str, k))
}
