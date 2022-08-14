package main

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	var maxLen int
	var subStart int
	for i:=0; i<len(s); i++ {
		length1 := expand(s, i, i)
		length2 := expand(s, i, i+1)

		length := max(length1, length2)
		if length > maxLen {
			maxLen = length
			subStart = i - (length-1)/2 // 回文前半部分
		}
	}

	return s[subStart:subStart+maxLen]
}

func expand(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// (right-1)-(left+1)+1
	return right-left-1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
