package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	var left int
	var right = len(nums) - 1
	for left <= right {
		mid := left + (right - left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left++
		} else {
			right--
		}
	}
	return -1
}
