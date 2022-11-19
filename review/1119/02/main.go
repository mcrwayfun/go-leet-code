package main

// time complexity: O(n)
// space complexity: O(1)
func isPalindrome(s string) bool {
	var start int
	var end = len(s) - 1
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

func isAlphanumeric(x byte) bool {
	return (x >= 'a' && x <= 'z') || (x >= 'A' && x <= 'Z') ||
		(x >= '0' && x <= '9')
}

func isEqualIgnoreCase(a, b byte) bool {
	if a >= 'A' && a <= 'Z' {
		a += 32
	}
	if b >= 'A' && b <= 'Z' {
		b += 32
	}
	return a == b
}
