package main

import "math"

// time complexity: O(n)
// space complexity: O(1)
func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	var n = len(s)
	var negative bool
	var p int
	for p < n && s[p] == ' ' {
		p++
	}
	if p == n { // 全部都是空格
		return 0
	}

	if p < n && s[p] == '+' { // 移除首位字符
		p++
	} else if p < n && s[p] == '-' {
		p++
		negative = true
	}

	for p < n && s[p] == '0' { // 移除前导0
		p++
	}

	// 开始计算数字个数
	var start = p
	for p < n && s[p] >= '0' && s[p] <= '9' {
		p++
	}
	if start == p { // 没有数字
		return 0
	}
	if p-start > 10 { // 超过极值，直接返回极值
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

	if num > math.MaxInt32 {
		if negative {
			return math.MinInt32
		} else {
			return math.MaxInt32
		}
	}

	if negative {
		return -int(num)
	}

	return int(num)
}
