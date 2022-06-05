package main

// time complexity: O(n)
// space complexity: O(1)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var duplicatedMap = make(map[byte]int, len(s))
	var maxLen int
	var start int
	for end:=0; end<len(s); end++ {
		if _, ok := duplicatedMap[s[end]]; ok {
			start = max(duplicatedMap[s[end]]+1, start)
		}
		duplicatedMap[s[end]] = end
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
