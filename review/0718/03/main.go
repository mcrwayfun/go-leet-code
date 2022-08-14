package main

// time complexity: O(m) // x和y的异或之后数字1的个数为m
// space complexity: O(1)
func hammingDistance(x int, y int) int {
	return countOne(x ^ y)
}

func countOne(n int) int {
	var count int
	for n != 0 {
		count++
		n &= n - 1
	}
	return count
}
