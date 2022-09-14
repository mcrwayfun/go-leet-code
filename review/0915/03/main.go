package main

func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}

	if n < 2 {
		return 1
	}

	prev2 := 1
	prev1 := 2
	for i := 2; i < n; i++ {
		cur := prev1 + prev2
		prev2 = prev1
		prev1 = cur
	}
	return prev1
}
