package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var duplicatedMap = make(map[byte]int, len(s))
	var start int
	var maxLen int
	for end := 0; end < len(s); end++ {
		if _, ok := duplicatedMap[s[end]]; ok {
			start = max(duplicatedMap[s[end]]+1, start)
		}
		duplicatedMap[s[end]] = end
		maxLen = max(end-start+1, maxLen)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
