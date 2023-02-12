package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var indexMap = make(map[byte]int, len(s))
	var start int
	var maxLen int
	for end := 0; end < len(s); end++ {
		if _, ok := indexMap[s[end]]; ok {
			start = max(indexMap[s[end]]+1, start)
		}
		indexMap[s[end]] = end
		maxLen = max(maxLen, end-start+1)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
