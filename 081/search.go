package _81

/**
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。

编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。

链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii
time complexity:O(lgn)
space complexity:O(1)
 */
func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		return nums[0] == target
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] > nums[left] { // 左边有序
			if nums[left] <= target && nums[mid] >= target {
				right = mid
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[left] { // 右边有序
			if nums[mid] <= target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid
			}
		} else {
			// 最坏的情况. 1 1 1 1 0 1 1
			left++
		}
	}

	// 最坏情况:1 1 1 1 1 1 1,等同于O(n)
	if nums[right] == target {
		return true
	}

	return false
}
