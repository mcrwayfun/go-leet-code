package main

// time complexity: O(n)
// space complexity: O(1)
func myPow(x float64, n int) float64 {
	var negative bool
	var N = n
	if n < 0 {
		negative = true
		N = -n
	}

	var res float64 = 1
	for i := 0; i < N; i++ {
		res *= x
	}

	if negative {
		return 1 / res
	}
	return res
}
