package main

func isPalindrome(x int) bool {
	var N = x
	if N <= 0 {
		return false
	}

	var res int
	for N != 0 {
		res = res * 10 + N % 10
		N /= 10
	}
	return res == x
}
