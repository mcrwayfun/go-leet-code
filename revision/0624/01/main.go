package main

import "math"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	var maxSum = math.MinInt64
	var cur int
	for i:=0; i<len(nums); i++ {
		if cur <= 0 {
			cur = nums[i]
		} else {
			cur += nums[i]
		}

		if cur > maxSum {
			maxSum = cur
		}
	}
	return maxSum
}
