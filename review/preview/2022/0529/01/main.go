package main

// time complexity: O(m)
// space complexity: O(1)
func myPow(x float64, n int) float64 {
	var res float64 = 1
	var N = abs(n)
	for N != 0 {
		if N & 1 == 1 {
			res *= x
		}
		x *= x
		N >>= 1
	}

	if n < 0 {
		return 1/res
	}

	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
