package main

// 使用遍历n的方式来达到求n次方
func myPow(x float64, n int) float64 {
	var N = abs(n)
	var result float64 = 1
	for i := 0; i < N; i++ {
		result *= x
	}

	if n < 0 {
		return 1 / result
	}
	return result
}

// 使用位运算的方式来解决
func myPowBitOperation(x float64, n int) float64 {
	var N = abs(n)
	var result float64 = 1

	for N != 0 {
		if N & 1 == 1 {
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

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
