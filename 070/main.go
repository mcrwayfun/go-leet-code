package main

import "fmt"

var m map[int]int

func climbStairs(n int) int {
	m = make(map[int]int, 0)
	return helper(n)
}

func helper(n int) int {
	if v, ok := m[n]; ok {
		return v
	}

	if n == 1 || n == 2{
		return n
	}

	steps := helper(n-1) + helper(n-2)
	m[n] = steps
	return steps
}

func main()  {
	steps := climbStairs(44)
	fmt.Println(steps)
}
