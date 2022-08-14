package main

// time complexity: O(n)
// space complexity: O(1)
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var curSum = nums[0]
	var maxSum = nums[0]
	for i:=1; i<len(nums); i++ {
		if curSum <= 0 {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
