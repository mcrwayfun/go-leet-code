package main

// time complexity: O(n)
// space complexity: O(1)
// 使用遍历n的方式，当n非常大的时候，会导致超时.
func myPowTimeOut(x float64, n int) float64 {
	var N = n
	if N < 0 {
		N = -N
	}

	var result float64 = 1
	for i := 0; i < N; i++ {
		result *= x
	}

	if n < 0 {
		return 1 / result
	}
	return result
}

// 另一种解题思路，遍历n的位数。n为int类型，在64位系统上最长为64位，
// 遍历n的位数，假设当前位数是a，当 a&1=1 需要将当前的结果累乘.
// time complexity: O(m)
// space complexity: O(1)
func myPow(x float64, n int) float64 {
	var N = n
	if N < 0 {
		N = -N
	}

	var result float64 = 1
	for N != 0 {
		if N&1 == 1 {
			result *= x
		}
		x *= x
		N >>= 1 // 将N不断向右移动，取高位来判断
	}

	if n < 0 {
		result = 1 / result
	}
	return result
}
