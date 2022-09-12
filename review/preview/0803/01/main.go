package main

import "fmt"

// time complexity: O(n)
// space complexity: O(1)
func longestPalindrome(s string) int {
	var charCountMap = make(map[byte]int, 256)
	for i := 0; i < len(s); i++ {
		charCountMap[s[i]]++
	}

	var oddCount int
	for _, count := range charCountMap {
		if count%2 != 0 {
			oddCount++
		}
	}

	unused := max(oddCount-1, 0)
	return len(s) - unused
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestPalindrome("abcba"))
}
