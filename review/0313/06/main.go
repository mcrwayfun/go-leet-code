package main

import "fmt"

func findPermutation(str string, pattern string) bool {
	if len(str) == 0 {
		return false
	}

	var charFrequencyMap = make(map[byte]int, len(pattern))
	for i := 0; i < len(pattern); i++ {
		charFrequencyMap[pattern[i]]++
	}

	var windowStart int
	var matched int

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]
		if _, ok := charFrequencyMap[rightChar]; ok {
			charFrequencyMap[rightChar]--
			if charFrequencyMap[rightChar] == 0 {
				matched++
			}
		}

		if matched == len(pattern) {
			return true
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

	return false
}

func main() {
	str := "oidbcaf"
	pattern := "abc"

	str2 := "odicf"
	pattern2 := "dc"

	str3 := "bcdxabcdy"
	pattern3 := "bcdyabcdx"

	str4 := "aaacb"
	pattern4 := "abc"

	fmt.Println(findPermutation(str, pattern))
	fmt.Println(findPermutation(str2, pattern2))
	fmt.Println(findPermutation(str3, pattern3))
	fmt.Println(findPermutation(str4, pattern4))
}
