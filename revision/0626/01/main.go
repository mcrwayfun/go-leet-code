package main

import "math"

/**
1：移除空格
2：判断正负
3：移除首位0
4：判断数字个数
5：没有数字则返回0
6：超过32位极值则返回极值
7：转换为num
8：超过32位极值则返回极值
9：根据正负返回对应的值
*/
func myAtoi(s string) int {
	var n = len(s)
	var start int
	var p int
	var negative bool

	for p < n && s[p] == ' ' {
		p++
	}
	if p == n {
		return 0
	}

	if s[p] == '+' {
		p++
	} else if s[p] == '-' {
		p++
		negative = true
	}

	for p < n && s[p] == '0' {
		p++
	}

	start = p
	for p < n && s[p] >= '0' && s[p] <= '9' {
		p++
	}
	if start == p {
		return 0
	}
	if p-start > 10 {
		if negative {
			return math.MinInt32
		}
		return math.MaxInt32
	}

	var num int64
	for i := start; i < p; i++ {
		num = num*10 + int64(s[i]-'0')
	}

	if num > math.MaxInt32 {
		if negative {
			return math.MinInt32
		}
		return math.MaxInt32
	}

	var res = int(num)
	if negative {
		return -res
	}
	return res
}
