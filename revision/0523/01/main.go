package main

import "fmt"

/**
使用遍历的方法来求值
*/
func myPow(x float64, n int) float64 {
	var N = abs(n)
	var result float64 = 1
	for i := 0; i < N; i++ {
		result *= x
	}

	if n < 0 {
		return 1/result
	}
	return result
}

func myPowBitOperation(x float64, n int) float64 {
	var N = abs(n)
	var result float64 = 1
	for N != 0 {
		if N & 1 == 1 {// 当前是有效的指数位
			result *= x
		}
		x *= x
		N >>= 1
	}

	if n < 0 {
		return 1/result
	}
	return result
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	fmt.Println(myPow(2.00000, 10))
	fmt.Println(myPowBitOperation(2.00000, 10))
}