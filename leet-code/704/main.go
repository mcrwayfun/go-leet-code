package main

/**
给定一个n个元素有序的（升序）整型数组nums 和一个目标值target ，
写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。

示例 1:
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4

示例2:
输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/binary-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func search(nums []int, target int) int {}
*/

/**
这是一道典型的二分法，通过不断缩小范围求得target==num，需要注意的是
1：mid=low+(high-low)/2, 避免循环过程中直接使用(high+low)/2溢出
2：for循环条件，需要注意数组中仅单个元素的情况，low<=high, 必须包含low==high

time complexity: O(lgn)
space complexity: O(1)
 */
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	var low int
	var high = len(nums) - 1
	for low <= high {
		mid := low + (high - low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
