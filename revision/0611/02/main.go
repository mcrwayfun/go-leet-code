package main

// time complexity: O(n)
// space complexity: O(1)
func climbStairs(n int) int {
	var first = 1
	var second = 1
	for i := 0; i < n; i++ {
		third := first + second
		first = second
		second = third
	}
	return first
}
