package _89_my_sqrt

func mySqrt(x int) int {
	left, right := 1, x+1
	for left < right {
		mid := left + (right-left)/2

		if mid*mid > x {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left - 1
}
