package main

import "fmt"

// time complexity: O(n)
// space complexity: O(n)
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	fn := make([]int, n+1)
	fn[1] = 1
	fn[2] = 2
	for i := 3; i <= n; i++ {
		fn[i] = fn[i-1] + fn[i-2]
	}
	return fn[n]
}

func main() {
	steps := climbStairs(2)
	fmt.Println(steps)

	steps = climbStairs(3)
	fmt.Println(steps)

	steps = climbStairs(6)
	fmt.Println(steps)

	steps = climbStairs(44)
	fmt.Println(steps)
}
