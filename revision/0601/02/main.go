package main

// time complexity: O(n)
// space complexity: O(1)
func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}

	var count int
	for i := 0; i < len(s); i++ {
		count += expand(s, i, i)
		count += expand(s, i, i+1)
	}
	return count
}

func expand(s string, left int, right int) int {
	var count int
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
		count++
	}
	return count
}
