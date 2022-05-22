package main

import (
	"fmt"
)

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

func myPowBitOperation(x float64, n int) float64 {
	var N = abs(n)
	var result float64 = 1
	for N != 0 {
		if N&1 == 1 { // 最后一位为1，表明当前存在有效指数
			result *= x
		}
		x *= x
		N = N >> 1
	}

	if n < 0 {
		return 1 / result
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(myPow(2.000000, 10))
	fmt.Println(myPowBitOperation(2.000000, 10))
}
