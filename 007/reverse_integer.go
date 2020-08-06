package _07

import (
	"math"
	"strconv"
)

// time complexity:O(n)
// space complexity:O(1)
func Reverse(x int) int {

	sign := 1
	if x < 0 {
		sign = -1
		x = -1 * x
	}
	// transform int to string
	str := strconv.Itoa(x)
	var result int64
	for i := len(str) - 1; i >= 0; i-- {
		// get the last digit
		temp := int64(x % 10)
		result = result*10 + temp
		x = x / 10
	}

	if result > math.MaxInt32 || result < -math.MaxInt32-1 {
		return 0
	}

	return int(result) * sign
}

func Reverse1(x int) int {

	sign := 1
	if x < 0 {
		sign = -1
		x = x * -1
	}

	var result int64
	for i := x; i != 0; i = i / 10 {
		result = result*10 + int64(i%10)
	}

	if result > math.MaxInt32 || result < -math.MaxInt32-1 {
		return 0
	}

	return int(result) * sign
}
