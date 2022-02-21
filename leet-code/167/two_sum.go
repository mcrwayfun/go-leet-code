package _67

import "sort"

/**
给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。

函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
说明:
	- 返回的下标值（index1 和 index2）不是从零开始的。
	- 你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。

链接：https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted

time complexity:O(n)
space complexity:O(n)
 */
func twoSum(numbers []int, target int) []int {
	sumMap := make(map[int]int)
	ret := make([]int, 0)

	for i, v := range numbers {
		sub := target - v
		if _, ok := sumMap[sub]; ok {
			ret = append(ret, i+1, sumMap[sub]+1)
			break
		}
		sumMap[v] = i
	}
	// 升序排序
	sort.Ints(ret)
	return ret
}

/**
有序
time complexity:O(lgn)
space complexity:O(1)
*/
func twoSum2(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}

	}
	return nil
}