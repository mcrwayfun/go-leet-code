package main

// time complexity: O(m)
// space complexity: O(1)
func myPow(x float64, n int) float64 {
	var N = n
	if n < 0 {
		N = -N
	}

	var result float64 = 1
	for N != 0 {
		if N&1 == 1 {
			result *= x
		}
		x *= x
		N >>= 1
	}
	if n < 0 {
		return 1 / result
	}
	return result
}
