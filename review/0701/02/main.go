package main

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	var start int
	var end = len(s)-1
	for start <= end {
		for start < end && !isAlphanumeric(s[start]) {
			start++
		}
		for start < end && !isAlphanumeric(s[end]) {
			end--
		}
		if !isEqual(s[start], s[end]) {
			return false
		}
		start++
		end--
	}
	return true
}

func isAlphanumeric(a byte) bool {
	return (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9')
}

func isEqual(a, b byte) bool {
	if a >= 'A' && a <= 'Z' {
		a += 32
	}
	if b >= 'A' && b <= 'Z' {
		b += 32
	}
	return a == b
}
