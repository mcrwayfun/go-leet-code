package main

import "fmt"

func findLength(str string) int {
	if len(str) == 0 {
		return 0
	}

	var start int
	var maxLength int
	var duplicatedMap = make(map[byte]int, len(str))
	for end := 0; end < len(str); end++ {
		if _, ok := duplicatedMap[str[end]]; ok {// shrink window
			start = duplicatedMap[str[end]] + 1
		}

		duplicatedMap[str[end]] = end
		if end - start +1 > maxLength {
			maxLength = end - start + 1
		}
	}

	return maxLength
}

func main() {
	str := "aabccbb"
	fmt.Println(findLength(str))
}