package main

/**
给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。

请你找出并返回只出现一次的那个数。

你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。

示例 1:
输入: nums = [1,1,2,3,3,4,4,8,8]
输出: 2

示例 2:
输入: nums =  [3,3,7,7,10,11,11]
输出: 10


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/single-element-in-a-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func singleNonDuplicate(nums []int) int {}
*/

/**
解法1：使用异或的方式可以求得单身数字
time complexity: O(n)
space complexity: O(1)
*/
func singleNonDuplicate(nums []int) int {
	var result = 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}

/**
注意题目有一个特性没有用上，即排序。可以使用二分法的方式来求得答案。
假设现在输入的数组是：[1,1,2,3,3,4,4,8,8]，非单身数字和它左右两边必定存在一个和它相等的数字，
可能会存在三种情况：
1：假设mid=6, nums[mid]=4, 它的左边和它相等，mid--
2：假设mid=5, nums[mid]=3, 它的右边和它相等，不做任何操作
3：否则只剩下一种情况，此时nums[mid]是单身数字，可以返回true

如果上述情况没有提早返回，则表示还没有找到单身数字：
1：如果(mid-low)%2 == 1, 表示左半边是奇数，
则单身数字肯定在左边（因为非单身是成对出现），mid = high - 1
2：否则 mid = low + 2
*/
func singleNonDuplicateBinarySearch(nums []int) int {
	var low = 0
	var high = len(nums) - 1

	for low <= high {// 如果数组只有一个元素0，low<=high 也成立
		mid := low + (high-low)/2
		if mid-1 >= 0 && nums[mid-1] == nums[mid] {
			mid--
		} else if mid+1 < len(nums) && nums[mid+1] == nums[mid] {
		} else {
			return nums[mid]
		}

		if (mid-low)%2 == 1 {
			high = mid - 1
		} else {
			low = mid + 2
		}
	}

	return -1
}
