package main

import "math"

// time complexity: O(n)
// space complexity: O(1)
func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	var n = len(s)
	var p int

	// 移除无用的空格
	for p < n && s[p] == ' ' {
		p++
	}
	if p == n {
		return 0
	}

	var negative bool
	if p < n && s[p] == '+' {
		p++
	} else if p < n && s[p] == '-' {
		p++
		negative = true
	}

	// 前导0忽略
	for p < n && s[p] == '0' {
		p++
	}

	// 判断数字个数
	var start = p
	for p < n && s[p] >= '0' && s[p] <= '9' {
		p++
	}
	if start == p { // 全是非数字
		return 0
	}

	// 超过极值
	if p-start > 10 {
		if negative {
			return math.MinInt32
		} else {
			return math.MaxInt32
		}
	}

	var num int64
	for i := start; i < p; i++ {
		num = num*10 + int64(s[i]-'0')
	}

	if negative {
		num = -num
	}

	if num < math.MinInt32 {
		return math.MinInt32
	} else if num > math.MaxInt32 {
		return math.MaxInt32
	}

	return int(num)
}
