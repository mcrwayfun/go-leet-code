package main

// time complexity: O(m)
// space complexity: O(1)
func getSum(a int, b int) int {
	for b != 0 {
		sum := a ^ b
		carry := a & b << 1

		a = sum
		b = carry
	}
	return a
}
