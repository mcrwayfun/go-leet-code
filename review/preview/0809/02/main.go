package main

// time complexity: O(n+m)
// space complexity: O(1)
func hammingDistance(x int, y int) int {
	return countOne(x ^ y)
}

func countOne(a int) int {
	var count int
	for a != 0 {
		count++
		a &= a - 1
	}
	return count
}
