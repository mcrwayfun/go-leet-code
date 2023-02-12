package main

// time complexity: O(m)
// space complexity: O(1)
func myPow(x float64, n int) float64 {
	var negative bool
	var N = n
	if N < 0 {
		negative = true
		N = -N
	}

	var res float64 = 1
	for N != 0 {
		if N & 1 == 1 {
			res *= x
		}
		x *= x
		N >>= 1
	}

	if negative {
		return 1/res
	}
	return res
}
