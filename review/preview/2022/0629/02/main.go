package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	var start int
	var end = len(nums) - 1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
