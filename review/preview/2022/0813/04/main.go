package main

// time complexity: O(m)
// space complexity: O(1)
func myPow(x float64, n int) float64 {
	var N = n
	if n < 0 {
		N = -N
	}

	var res float64 = 1
	for N != 0 {
		if N & 1 == 1 {
			res *= x
		}
		x *= x
		N >>= 1 // 使用进位的方式来求n方
	}

	if n < 0 {
		return 1/res
	}
	return res
}


