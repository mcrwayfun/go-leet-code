package main

import "fmt"

func findSubstring(str string, pattern string) string {
	if len(str) == 0 {
		return ""
	}

	var charFrequencyMap = make(map[byte]int, len(pattern))
	for _, v := range []byte(pattern) {
		charFrequencyMap[v]++
	}

	var start int
	var matched int
	var subStart int
	var minLength = len(str) + 1

	for end := 0; end < len(str); end++ {
		curChar := str[end]
		if _, ok := charFrequencyMap[curChar]; ok {
			charFrequencyMap[curChar]--
			if charFrequencyMap[curChar] >= 0 { // pattern存在重复的字符
				matched++
			}
		}

		for matched == len(pattern) {
			if end-start+1 < minLength {
				minLength = end - start + 1
				subStart = start
			}

			startChar := str[start]
			if _, ok := charFrequencyMap[startChar]; ok {
				if charFrequencyMap[startChar] == 0 {
					matched--
				}
				charFrequencyMap[startChar]++
			}
			start++
		}
	}

	if minLength == len(str)+1 {
		return ""
	}

	return str[subStart : subStart+minLength]
}

func main() {
	str := "aabdec"
	pattern := "abc"
	fmt.Println(findSubstring(str, pattern))

	str2 := "abdabca"
	pattern2 := "abc"
	fmt.Println(findSubstring(str2, pattern2))
}