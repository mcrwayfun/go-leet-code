package main

import "fmt"

func findLength(str string) int {
	if len(str) == 0 {
		return 0
	}

	var windowStart int
	var charIndexMap = make(map[byte]int, len(str))
	var maxLength int

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		char := str[windowEnd]
		if _, ok := charIndexMap[char]; ok {
			windowStart = charIndexMap[char] + 1
		}
		charIndexMap[char] = windowEnd
		maxLength = max(maxLength, windowEnd-windowStart+1)
	}

	return maxLength
}

func main() {
	str := "aabccbb"
	str2 := "abbbb"
	str3 := "abccde"

	fmt.Println(findLength(str))
	fmt.Println(findLength(str2))
	fmt.Println(findLength(str3))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
