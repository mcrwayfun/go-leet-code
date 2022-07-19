package main

// time complexity: O(n)
// space complexity: O(1)
func findKthLargest(nums []int, k int) int {
	var low = 0
	var high = len(nums) - 1
	for low <= high {
		var p = partition(nums, low, high)
		if p == k-1 {
			return nums[p]
		} else if p < k-1 {
			low = p + 1
		} else {
			high = p - 1
		}
	}
	return -1
}

func partition(nums []int, low, high int) int {
	var pivot = nums[low]
	var i = low
	var j = high

	for i < j {
		for ; i < j && nums[j] < pivot; j-- {
		}
		nums[i], nums[j] = nums[j], nums[i]
		for ; i < j && nums[i] >= pivot; i++ {
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return i
}
