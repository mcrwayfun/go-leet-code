package _07_reverse_integer

import (
	"math"
	"strconv"
)

func Reverse(x int) int {

	var sign int
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
