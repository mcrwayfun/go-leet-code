package main

import "math"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	var sum int
	var maxSum = math.MinInt64
	for i:=0; i<len(nums); i++ {
		if sum <= 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}

		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}
