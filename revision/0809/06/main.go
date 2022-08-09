package main

// time complexity: O(n*lgn)
// space complexity: O(1)
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return -1
	}

	var low int
	var high = len(nums) - 1
	for low <= high {
		p := partition(nums, low, high)
		if p == k - 1 {
			return nums[p]
		} else if p < k - 1 {
			low = p + 1
		} else {
			high = p - 1
		}
	}
	return -1
}

func partition(nums []int, low, high int) int {
	var pivot = nums[low]
	for low < high {
		for ; low < high && nums[high] < pivot ; high--{}
		nums[low], nums[high] = nums[high], nums[low]
		for ; low < high && nums[low] >= pivot ; low++{}
		nums[low], nums[high] = nums[high], nums[low]
	}
	return low
}
