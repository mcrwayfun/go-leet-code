package main

// time complexity: O(n^2)
// space complexity: O(1)
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	var maxLength int
	var subStart int
	for i := 0; i < len(s); i++ {
		length1 := expand(s, i, i)
		length2 := expand(s, i, i+1)
		length := max(length1, length2)

		if length > maxLength {
			maxLength = length
			subStart = i - (length-1)/2
		}
	}
	return s[subStart : subStart+maxLength]
}

func expand(s string, start, end int) int {
	for start >= 0 && end < len(s) && s[start] == s[end] {
		start--
		end++
	}
	// end - 1 - (start + 1 ) + 1 =
	return end - start - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
