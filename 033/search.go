package main

import "fmt"

/**
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
time complexity:O(lgn)
space complexity:O(1)
*/
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 && nums[0] == target {
		return 0
	}

	// 临界值的控制. 当right = len(nums)-1,left<=right
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < nums[right] { // 右边有序
			// target在右边
			if nums[mid] <= target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid
			}
		} else {
			// 左边有序
			// target在左边
			if nums[mid] >= target && nums[left] <= target {
				right = mid
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}

func search2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		} else if nums[start] <= nums[mid] {
			if nums[start] <= target && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{3, 1}
	target := 1
	index := search2(nums, target)
	fmt.Println(index)
}
