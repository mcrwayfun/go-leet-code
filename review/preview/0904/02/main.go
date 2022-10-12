package main

// time complexity: O(n)
// space complexity: O(n)
func longestPalindrome(s string) int {
	if len(s) == 0 {
		return 0
	}

	var charFrequencyMap = make(map[byte]int, len(s))
	for i := range s {
		charFrequencyMap[s[i]]++
	}

	var oddCount int
	for _, count := range charFrequencyMap {
		if count%2 == 1 {
			oddCount++
		}
	}

	return len(s) - max(oddCount-1, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
