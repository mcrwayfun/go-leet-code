package main

// time complexity: O(n*lgn)
// space complexity: O(1)
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 || k > len(nums) {
		return -1
	}

	var start int
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
	var pivot = nums[start]
	for start < end {
		for ; start < end && nums[end] <= pivot; end-- {
		}
		nums[start], nums[end] = nums[end], nums[start]
		for ; start < end && nums[start] > pivot; start++ {
		}
		nums[start], nums[end] = nums[end], nums[start]
	}
	return start // start == end
}
