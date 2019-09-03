package _34_search_range

/**
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
time complexity:O(lgn)
space complexity:O(1)
*/
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	// 第一个出现的位置
	firstPosition := findTarget(nums, target)
	// 超过数组长度或者找不到target
	if firstPosition == len(nums) || nums[firstPosition] != target {
		return []int{-1, -1}
	}
	// 使用target+1去遍历,存在两种情况（即使target+1出现的位置越界，减1后也不会越界）
	// 1:target+1存在,那么它出现的位置减1必定是target出现的最后一个位置
	// 2:target+1不存在,那么返回的结果减1也必定是target出现的最后一个位置
	lastPosition := findTarget(nums, target+1) - 1

	return []int{firstPosition, lastPosition}
}

func findTarget(nums []int, target int) int {
	left, right := 0, len(nums)

	// 找到第一个target出现的位置,所以不用判断 nums[mid] == target的情况
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
