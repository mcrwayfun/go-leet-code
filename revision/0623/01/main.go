package main

import "math"

func myAtoi(s string) int {
	var start int
	var p int
	var negative bool

	// remove space
	for p < len(s) && s[p] == ' ' {
		p++
	}

	if s[p] == '+' {
		p++
	} else if s[p] == '-' {
		p++
		negative = true
	}

	if p < len(s) && s[p] == '0' {
		p++
	}

	start = p
	for p < len(s) && s[p] >= '0' && s[p] <= '9' {
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
	for i:=start; i<len(s); i++ {
		num = num * 10 + int64(s[i] - '0')
	}

	if num > math.MaxInt32 {
		if negative {
			return math.MinInt32
		}
		return math.MaxInt32
	}

	res := int(num)
	if negative {
		return -res
	}

	return res
}

func main() {
	myAtoi("4193 with words")
}
