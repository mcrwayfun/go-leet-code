package main

func hammingDistance(x int, y int) int {
	return countNumberOfOne(x ^ y)
}

func countNumberOfOne(n int) int {
	var count int
	for n != 0 {
		count++
		n &= n - 1
	}
	return count
}
