package main

func myPow(x float64, n int) float64 {
	var N = n
	if N < 0 {
		N = -N
	}
	var ret float64 = 1
	for N != 0 {
		if N&1 == 1 {
			ret *= x
		}
		x *= x
		N >>= 1
	}

	if n < 0 {
		return 1 / ret
	}
	return ret
}
