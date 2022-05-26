package main

import "math"

/**
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组 是数组中的一个连续部分。

示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组[4,-1,2,1] 的和最大，为6 。

示例 2：
输入：nums = [1]
输出：1

示例 3：
输入：nums = [5,4,-1,7,8]
输出：23

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func maxSubArray(nums []int) int
*/

/**
1：维护两个变量maxSum和cur(当前窗口的和)
2：遍历数组，将当前元素值加到cur中
	2-1：发现当cur<=0时，对maxSum无正向影响，直接取当前的ele作为cur
	2-2：否则 cur=cur+ele

time complexity: O(n)
space complexity: O(1)
*/
func maxSubArray(nums []int) int {
	var maxSum = math.MinInt64
	var cur int
	for i := 0; i < len(nums); i++ {
		if cur <= 0 {
			cur = nums[i]
		} else {
			cur += nums[i]
		}

		if maxSum < cur {
			maxSum = cur
		}
	}
	return maxSum
}
