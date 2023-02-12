package main

// time complexity: O(n)
// space complexity: O(1)
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 || len(nums) < k {
		return -1
	}

	var start = 0
	var end = len(nums) - 1
	for start <= end {
		p := partition(nums, start, end)
		if p == k-1 {
			return nums[p]
		} else if p < k-1 {
			start = p + 1
		} else {
			end = p - 1
		}
	}
	return -1
}

func partition(nums []int, start, end int) int {
	pivot := nums[start]
	for start < end {
		for ; start < end && nums[end] <= pivot; end-- {}
		nums[start], nums[end] = nums[end], nums[start]
		for ; start < end && nums[start] >= pivot; start++{}
		nums[start], nums[end] = nums[end], nums[start]
	}
	return start
}
