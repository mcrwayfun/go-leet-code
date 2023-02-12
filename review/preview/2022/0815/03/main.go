package main

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	var N = n
	for N % 5 == 0 {
		N /= 5
	}
	for N % 3 == 0 {
		N /= 3
	}
	for N % 2 == 0 {
		N /= 2
	}
	return N == 1
}
