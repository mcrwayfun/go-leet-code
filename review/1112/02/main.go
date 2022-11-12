package main

import "math"

func myAtoi(s string) int {
	var n = len(s)
	var p int

	// 读入字符串并丢弃无用的前导空格
	for p < n && s[p] == ' ' {
		p++
	}
	if p == n {
		return 0
	}

	// 检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符
	var negative bool
	if s[p] == '-' {
		negative = true
		p++
	} else if s[p] == '+' {
		p++
	}

	// 忽略前导0
	for p < n && s[p] == '0' {
		p++
	}

	// 读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
	var q = p
	for p < n && s[p] >= '0' && s[p] <= '9' {
		p++
	}
	if q == p {
		return 0 // 没有数字
	}

	if p-q > 10 { // 超出极值
		if negative {
			return math.MinInt32
		}
		return math.MaxInt32
	}

	var num int
	for i := q; i < p; i++ {
		num = num*10 + int(s[i]-'0')
	}

	if negative {
		num = -num
	}

	if num < math.MinInt32 {
		return math.MinInt32
	} else if num > math.MaxInt32 {
		return math.MaxInt32
	}

	return num
}
