package main

func climbStairs(n int) int {
	first := 1
	second := 1
	for i:=1; i<n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}
