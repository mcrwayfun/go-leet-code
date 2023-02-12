package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	var n = len(nums)
	return max(helper(nums, 0, n-2), helper(nums, 1, n-1))
}

func helper(nums []int, start, end int) int {
	var prev1 int
	var prev2 int
	for i := start; i <= end; i++ {
		cur := max(prev1, prev2+nums[i])
		prev2 = prev1
		prev1 = cur
	}
	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}