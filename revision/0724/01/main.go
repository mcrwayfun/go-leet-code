package main

// time complexity: O(n)
// space complexity: O(n)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var noDuplicatedMap = make(map[byte]int)
	var start int
	var maxLength int
	for end := 0; end < len(s); end++ {
		if _, ok := noDuplicatedMap[s[end]]; ok {
			start = max(noDuplicatedMap[s[end]]+1, start)
		}

		maxLength = max(maxLength, end-start+1)
		noDuplicatedMap[s[end]] = end
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
