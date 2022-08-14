package main

// time complexity: O(n^2)
// space complexity: O(1)
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	var maxlen int
	var substart int
	for i:=0; i<len(s); i++ {
		length1 := expand(s, i, i)
		length2 := expand(s, i, i+1)
		length := max(length1, length2)
		if length > maxlen {
			maxlen = length
			substart = i - (length-1)/2 // 回文的前半部分
		}
	}
	return s[substart:substart+maxlen]
}

func expand(s string, left int, right int) int {
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
