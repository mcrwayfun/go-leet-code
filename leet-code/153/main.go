package main

import (
	"fmt"
)

/// Time Complexity: O(lgn)
/// Space Complexity: O(1)
func findMin(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		if nums[start] < nums[end] {
			return nums[start]
		}

		mid := start + (end-start)/2
		if nums[mid] > nums[end] {
			start = mid
		} else {
			end = mid
		}
	}
	return min(nums[start], nums[end])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	nums := []int{11, 13, 15, 17}
	min := findMin(nums)
	fmt.Println(min)
}
