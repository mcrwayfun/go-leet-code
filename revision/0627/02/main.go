package main

func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}

	var res int
	for i:=0; i<len(s); i++ {
		res += expand(s, i, i)
		res += expand(s, i, i+1)
	}
	return res
}

func expand(s string, left, right int) int {
	var count int
	for left >= 0 && right < len(s) && s[left] == s[right] {
		count++
		left--
		right++
	}
	return count
}
