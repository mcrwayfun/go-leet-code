package main

import "math"

// time complexity: O(n)
// space complexity: O(1)
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	var maxSum = math.MinInt64
	var cur int
	for _, num := range nums {
		if cur <= 0 {
			cur = num
		} else {
			cur += num
		}
		maxSum = max(maxSum, cur)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
