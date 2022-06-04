package main

// time complexity: O(n)
// space complexity: O(1)
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	var left, right = 0, len(s) - 1
	for left < right {
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}
		if !isEqualIgnoreCase(s[left], s[right]) {
			return false
		}
		left++
		right--
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
