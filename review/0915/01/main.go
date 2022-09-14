package main

func hammingWeight(num uint32) int {
	var N = num
	var count int
	for N != 0 {
		count++
		N = N & (N - 1)
	}
	return count
}
