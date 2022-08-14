package main

import "strconv"

/**
1：负数肯定不是回文数字
2：假设现在有数字x，x%10可以得到尾数，x/10可以减少尾数，申明变量sum。
	2-1: sum += sum*10 + x%10
	2-2: x =x/10
time complexity: O(m)
space complexity: O(1)
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var tmp = x
	var y int
	for tmp != 0 {
		y = y*10 + tmp%10
		tmp /= 10
	}
	return x == y
}

func isPalindromeStr(x int) bool {
	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)
	var left, right = 0, len(str) - 1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
