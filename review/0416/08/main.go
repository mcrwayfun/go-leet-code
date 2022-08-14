package main

import "fmt"

// 给一个字符串和pattern，找到所有符合pattern的排列组合, 返回起始的下标
// str=ppqp, pattern=pq, ret=[1, 2]
// if windowEnd-windowStart+1 == len(pattern) 则滑动窗口
func findStringAnagrams(str string, pattern string) []int {
	if len(str) == 0 || len(pattern) == 0 {
		return []int{-1, -1}
	}

	var charFrequencyMap = make(map[byte]int, len(pattern))
	for i := 0; i < len(pattern); i++ {
		charFrequencyMap[pattern[i]]++
	}

	var windowStart int
	var startIndexes []int
	var matched int

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		curChar := str[windowEnd]
		if _, ok := charFrequencyMap[curChar]; ok {
			charFrequencyMap[curChar]--
			if charFrequencyMap[curChar] == 0 {
				matched++
			}
		}

		if matched == len(charFrequencyMap) {
			startIndexes = append(startIndexes, windowStart)
		}

		if windowEnd-windowStart+1 == len(pattern) {
			startChar := str[windowStart]
			if _, ok := charFrequencyMap[startChar]; ok {
				if charFrequencyMap[startChar] == 0 {
					matched--
				}
				charFrequencyMap[startChar]++
			}
			windowStart++
		}
	}

	return startIndexes
}

func main() {
	// str=ppqp, pattern=pq, ret=[1, 2]
	str := "ppqp"
	pattern := "pq"
	fmt.Println(findStringAnagrams(str, pattern))

	str = "abbcabc"
	pattern = "abc"
	fmt.Println(findStringAnagrams(str, pattern))
}
