package main

import "fmt"

func findPermutation(str string, pattern string) bool {
	if len(str) == 0 || len(pattern) == 0 {
		return false
	}

	var charFrequencyMap = make(map[byte]int, len(pattern))
	for i := 0; i < len(pattern); i++ {
		charFrequencyMap[pattern[i]]++
	}

	var windowStart int
	var matched int
	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		character := str[windowEnd]
		if _, ok := charFrequencyMap[character]; ok {
			charFrequencyMap[character]--
			if charFrequencyMap[character] == 0 {
				matched++
			}
		}

		if len(pattern) == matched {
			return true
		}

		if windowEnd - windowStart + 1 >= len(pattern) {// shrink window
			startChar := str[windowStart]
			if _, ok:= charFrequencyMap[startChar]; ok {
				if charFrequencyMap[startChar] == 0 {
					matched--
				}
				charFrequencyMap[startChar]++
			}
			windowStart++
		}
	}

	return false
}

func main() {
	str := "oidbcaf"
	pattern := "abc"
	fmt.Println(findPermutation(str, pattern))

	str = "odicf"
	pattern = "dc"
	fmt.Println(findPermutation(str, pattern))

	str = "bcdxabcdy"
	pattern = "bcdyabcdx"
	fmt.Println(findPermutation(str, pattern))

	str = "aaacb"
	pattern = "abc"
	fmt.Println(findPermutation(str, pattern))
}