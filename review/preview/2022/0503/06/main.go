package main

// time complexity: O(n)
// space complexity: O(1)
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	start, end := 0, len(s)-1
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

// 判断是否为字母或者数字
func isAlphanumeric(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') ||
		(b >= '0' && b <= '9')
}

// 判断两个字符是否相等
func isEqualIgnoreCase(a, b byte) bool {
	if a >= 'A' && a <= 'Z' {
		a += 32
	}
	if b >= 'A' && b <= 'Z' {
		b += 32
	}
	return a == b
}
