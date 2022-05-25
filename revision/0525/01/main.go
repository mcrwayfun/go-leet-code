package main

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	start, end := 0, len(s) -1
	for start < end {
		for start < end && !isAlphanumeric(s[start]) {
			start++
		}
		for start < end && !isAlphanumeric(s[end]) {
			end--
		}
		if start < end && !isEqualIgnoreCase(s[start], s[end]) {
			return false
		}
		start++
		end--
	}
	return true
}

func isAlphanumeric(b byte) bool {
	return (b <= 'Z' && b >= 'A') || (b <= 'z' && b >= 'a') || (b <= '9' && b >= '0')
}

func isEqualIgnoreCase(a, b byte) bool {
	if a <= 'Z' && a >= 'A' {
		a += 32
	}
	if b <= 'Z' && b >= 'A' {
		b += 32
	}
	return a == b
}