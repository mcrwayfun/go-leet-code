package main

// time complexity: O(lgn)
// space complexity: O(1)
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	var left int
	var right = len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
