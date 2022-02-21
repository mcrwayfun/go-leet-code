package main

import "fmt"

func mySqrt(x int) int {
	start, end, ans := 0, x, -1
	for start <= end {
		mid := start + (end-start)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			start = mid + 1
			ans = mid
		} else if mid*mid > x {
			end = mid - 1
		}
	}

	return ans
}

func main() {
	sqrt := mySqrt(1)
	fmt.Println(sqrt)
}
