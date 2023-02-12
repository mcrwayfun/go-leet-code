package main

/**
1：如果是非字母或者数字，则需要忽略且跳过
2：使用双指针start和end，分别从字符串的起始和结束一一对应判断是否相等
*/
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true // 空字符肯定是回文
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

// 判断是否为字母和数字
func isAlphanumeric(b byte) bool {
	return (b <= 'Z' && b >= 'A') || (b <= 'z' && b >= 'a') || (b <= '9' && b >= '0')
}

// 判断a和b是否相等
func isEqualIgnoreCase(a, b byte) bool {
	if a <= 'Z' && a >= 'A' {
		a += 32
	}
	if b <= 'Z' && b >= 'A' {
		b += 32
	}
	return a == b
}
