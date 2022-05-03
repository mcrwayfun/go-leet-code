package main

/**
给你一个下标从 1 开始的整数数组numbers ，该数组已按 非递减顺序排列 ，请你从数组中找出满足相加之和等于目标数target 的两个数。
如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。

以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。

你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。

你所设计的解决方案必须只使用常量级的额外空间。


示例 1：

输入：numbers = [2,7,11,15], target = 9
输出：[1,2]
解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。
示例 2：

输入：numbers = [2,3,4], target = 6
输出：[1,3]
解释：2 与 4 之和等于目标数 6 。因此 index1 = 1, index2 = 3 。返回 [1, 3] 。
示例 3：

输入：numbers = [-1,0], target = -1
输出：[1,2]
解释：-1 与 0 之和等于目标数 -1 。因此 index1 = 1, index2 = 2 。返回 [1, 2]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func twoSum(numbers []int, target int) []int
*/

/**
题目解析：这是两数之和的变种题目，有一点不同的是：两数之和是无序数组，而这道题已经是按照递增排好序的。
所以可以不使用map来记录已经访问过的元素，改用双指针的方式来同时访问两个元素实现求和的目的。

解法：
1：维护两个指针start和end，start从数组下标0开始，而end从数组末尾开始
2：开始循环(start<end)
	2-1: 如果nums[start]+nums[end]==target, return {start+1,end+1}
	2-2: 如果nums[start]+nums[end]<target, start++
	2-3: 如果nums[start]+nums[end]>target, end--
3: 结束循环仍未找到，return {-1,-1}

Time Complexity: O(n)
Space Complexity: O(1)
*/
func twoSum(numbers []int, target int) []int {
	if len(numbers) == 0 {
		return []int{-1, -1}
	}

	var start int
	var end = len(numbers) - 1
	for start < end {
		sum := numbers[start] + numbers[end]
		if sum == target {
			return []int{start + 1, end + 1}
		} else if sum < target {
			start++
		} else {
			end--
		}
	}
	return []int{-1, -1}
}
