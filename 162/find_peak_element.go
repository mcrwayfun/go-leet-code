package _62

/**
峰值元素是指其值大于左右相邻值的元素。

给定一个输入数组 nums，其中 nums[i] ≠ nums[i+1]，找到峰值元素并返回其索引。

数组可能包含多个峰值，在这种情况下，返回任何一个峰值所在位置即可。

你可以假设 nums[-1] = nums[n] = -∞。

链接：https://leetcode-cn.com/problems/find-peak-element
time complexity:O(lgn)
space complexity:O(1)
*/
func findPeakElement(nums []int) int {
	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] < nums[mid+1] {
			start = mid
		} else {
			end = mid
		}
	}

	if nums[start] > nums[end] {
		return start
	} else {
		return end
	}
}
