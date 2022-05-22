package main

import "fmt"

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	var start = 0
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

func isEqualIgnoreCase(a, b byte) bool {
	if a <= 'Z' && a >= 'A' {
		a += 32
	}
	if b <= 'Z' && b >= 'A' {
		b += 32
	}
	return a == b
}

// 判断是不是字母或者数字
func isAlphanumeric(b byte) bool {
	return (b <= 'z' && b >= 'a') || (b <= 'Z' && b >= 'A') || (b <= '9' && b >= '0')
}

// docs
// docs update
func main() {
	fmt.Println(isPalindrome("race a car"))
}
