package main

/// Time Complexity: O(lgn)
/// Space Complexity: O(1)
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			right = mid// 左闭右开
		}else if nums[mid] > target {
			right = mid
		}else if nums[mid] < target {
			left = mid + 1
		}
	}

	return left
}