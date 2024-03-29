package main

// time complexity: O(m)
// space complexity: O(1)
func hammingWeight(num uint32) int {
	var count int
	for num != 0 {
		count++
		num &= num - 1
	}
	return count
}
