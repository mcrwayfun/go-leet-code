package main

import "fmt"

func findLength(str string, k int) int {
	if len(str) == 0 {
		return 0
	}

	var charFrequencyMap = make(map[byte]int, len(str))
	var start int
	var maxLength int

	for end := 0; end < len(str); end++ {
		charFrequencyMap[str[end]]++
		for len(charFrequencyMap) > k {
			if _, ok := charFrequencyMap[str[start]]; ok {
				charFrequencyMap[str[start]]--
				if charFrequencyMap[str[start]] == 0 {
					delete(charFrequencyMap, str[start])
				}
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
	str := "araaci"
	fmt.Println(findLength(str, 2))
}
