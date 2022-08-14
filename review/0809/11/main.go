package main

// time complexity: O(n)
// space complexity: O(1)
func longestPalindrome(s string) int {
	var countMap = make(map[byte]int, 256)
	for i := 0; i < len(s); i++ {
		countMap[s[i]]++
	}

	var oddCount int
	for _, count := range countMap {
		if count%2 == 1 {
			oddCount++
		}
	}

	var unused = max(oddCount-1, 0)
	return len(s) - unused
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
