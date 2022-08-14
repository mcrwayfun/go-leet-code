package main

import "fmt"

// 维护窗口中出现最多的字符 出现的次数
func findLength(str string, k int) int {
	if len(str) == 0 {
		return 0
	}

	var start int
	var maxRepeatCount int
	var maxLength int
	var charFrequencyMap = make(map[byte]int, len(str))

	for end := 0; end < len(str); end++ {
		charFrequencyMap[str[end]]++
		if charFrequencyMap[str[end]] > maxRepeatCount {
			maxRepeatCount = charFrequencyMap[str[end]]
		}

		if end-start+1-maxRepeatCount > k {
			if _, ok := charFrequencyMap[str[start]]; ok {
				charFrequencyMap[str[start]]--
			}
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

	str2 := "aaaabbbaaacd"
	fmt.Println(findLength(str, k))
	fmt.Println(findLength(str2, k))
}
