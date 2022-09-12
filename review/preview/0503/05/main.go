package main

import "fmt"

func findLength(str string, k int) int {
	if len(str) == 0 {
		return 0
	}

	var start int
	var maxRepeatCount int
	var duplicatedMap = make(map[byte]int, len(str))
	var maxLength int

	for end := 0; end < len(str); end++ {
		endChar := str[end]
		duplicatedMap[endChar]++
		if duplicatedMap[endChar] > maxRepeatCount {
			maxRepeatCount = duplicatedMap[endChar]
		}

		if end-start+1-maxRepeatCount > k {
			startChar := str[start]
			duplicatedMap[startChar]--
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

	str2 := "abbcb"
	k2 := 1
	fmt.Println(findLength(str2, k2))
}
