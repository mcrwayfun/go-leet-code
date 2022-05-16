package main

import "fmt"

func findLength(str string, k int) int {
	if len(str) == 0 {
		return 0
	}

	var start int
	var charFrequencyMap = make(map[byte]int, len(str))
	var maxLength int

	for end := 0; end < len(str); end++ {
		charFrequencyMap[str[end]]++

		for len(charFrequencyMap) > k {
			charFrequencyMap[str[start]]--
			if charFrequencyMap[str[start]] == 0 {
				delete(charFrequencyMap, str[start])
			}
			start++
		}

		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
	}

	return maxLength
}

func main() {
	fmt.Println(findLength("araaci", 2))
	fmt.Println(findLength("araaci", 1))
	fmt.Println(findLength("cbbebi", 3))
}
