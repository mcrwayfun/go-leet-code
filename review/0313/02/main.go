package main

import "fmt"

func findSubstring(str string, pattern string) string {
	if len(str) == 0 {
		return ""
	}

	var charFrequencyMap = make(map[byte]int, len(pattern))
	for i := 0; i < len(pattern); i++ {
		charFrequencyMap[pattern[i]]++
	}

	var windowStart int
	var matched int
	var minLength = len(str) + 1
	var subStart int

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]
		if _, ok := charFrequencyMap[rightChar]; ok {
			charFrequencyMap[rightChar]--
			if charFrequencyMap[rightChar] >= 0 {
				matched++
			}
		}

		for matched == len(pattern) {
			if minLength > windowEnd-windowStart+1 {
				minLength = windowEnd - windowStart + 1
				subStart = windowStart
			}

			leftChar := str[windowStart]
			if _, ok := charFrequencyMap[leftChar]; ok {
				if charFrequencyMap[leftChar] == 0 {
					matched--
				}
				charFrequencyMap[leftChar]++
			}
			windowStart++
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

	str2 := "abdabca"
	pattern2 := "abc"

	str3 := "adcad"
	pattern3 := "abc"

	fmt.Println(findSubstring(str, pattern))
	fmt.Println(findSubstring(str2, pattern2))
	fmt.Println(findSubstring(str3, pattern3))
}
