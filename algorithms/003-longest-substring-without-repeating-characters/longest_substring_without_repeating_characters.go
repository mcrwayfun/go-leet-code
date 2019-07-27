package _03_longest_substring_without_repeating_characters

import "strings"

// time complexity:O(n)
// space complexity:O(min(n,m))
func lengthOfLongestSubstring(s string) int {
	// s is empty
	if len(s) == 0 {
		return 0
	}
	// use map to store appear str
	tempMap := make(map[byte]int)
	left, maxCount := 0, 0
	for i := 0; i < len(s); i++ {
		if _, ok := tempMap[s[i]]; ok {
			left = max(left, tempMap[s[i]]+1)
		}
		tempMap[s[i]] = i
		maxCount = max(maxCount, i-left+1)
	}

	return maxCount
}

// time complexity:O(n)
// space complexity:O(1)
func lengthOfLongestSubstring1(s string) int {
	// s is empty
	if len(s) == 0 {
		return 0
	}
	left, maxCount := 0, 0
	for i := 0; i < len(s); i++ {
		if n := strings.Index(s[left:i], string(s[i])); n >= 0 {
			// slide the left
			left += n + 1
		}

		if i-left+1 > maxCount {
			maxCount = i - left + 1
		}
	}

	return maxCount
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
