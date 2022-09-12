package main

import "math"

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	var start int
	var p int
	var n = len(s)
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

	if p - start > 10 {
		if negative {
			return math.MinInt32
		}
		return math.MaxInt32
	}

	var num int64
	for i:=start; i<p; i++ {
		num = num * 10 + int64(s[i] - '0')
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
