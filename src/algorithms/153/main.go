package main

import (
	"fmt"
)

/// Time Complexity: O(lgn)
/// Space Complexity: O(1)
func findMin(nums []int) int {
	min := 1 << 31
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {
	nums := []int{3, 4, 5, 1, 2}
	min := findMin(nums)
	fmt.Println(min)
}
