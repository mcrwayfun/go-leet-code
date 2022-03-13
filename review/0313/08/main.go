package main

import "fmt"

func findStringAnagrams(str string, pattern string) []int {
	if len(str) == 0 {
		return []int{}
	}

	var charFrequencyMap = make(map[byte]int, len(pattern))
	for i := 0; i < len(pattern); i++ {
		charFrequencyMap[pattern[i]]++
	}

	var windowStart int
	var matched int
	var startIndices []int

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]
		if _, ok := charFrequencyMap[rightChar]; ok {
			charFrequencyMap[rightChar]--
			if charFrequencyMap[rightChar] == 0 {
				matched++
			}
		}

		if matched == len(pattern) {
			startIndices = append(startIndices, windowStart)
		}

		if windowEnd-windowStart+1 >= len(pattern) {
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

	return startIndices
}

func main() {
	str := "ppqp"
	pattern := "pq"

	str2 := "abbcabc"
	pattern2 := "abc"

	fmt.Println(findStringAnagrams(str, pattern))
	fmt.Println(findStringAnagrams(str2, pattern2))
}
