package main

import "fmt"

func findLength(str string) int {
	if len(str) == 0 {
		return 0
	}

	var duplicatedMap = make(map[byte]int, len(str))
	var start int
	var maxLength int

	for end := 0; end < len(str); end++ {
		endChar := str[end]
		if _, ok := duplicatedMap[endChar]; ok {
			start = duplicatedMap[endChar] + 1
		}
		duplicatedMap[endChar] = end // 存放的是下标

		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
	}

	return maxLength
}

func main() {
	str := "aabccbb"
	fmt.Println(findLength(str))
}
