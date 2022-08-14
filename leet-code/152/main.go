package main

/**
给你一个整数数组 nums，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

测试用例的答案是一个32-位 整数。

子数组 是数组的连续子序列。

示例 1:
输入: nums = [2,3,-2,4]
输出: 6
解释:子数组 [2,3] 有最大乘积 6。

示例 2:
输入: nums = [-2,0,-1]
输出: 0
解释:结果不能为 2, 因为 [-2,-1] 不是子数组。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-product-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func maxProduct(nums []int) int {}
*/

// time complexity: O(n)
// space complexity: O(1)
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var maxSum = nums[0]
	var curMax = nums[0]
	var curMin = nums[0]
	for i := 0; i < len(nums); i++ {
		a := curMax * nums[i]
		b := curMin * nums[i]
		c := nums[i]

		curMax = max3(a, b, c)
		curMin = min3(a, b, c)
		maxSum = max(curMax, maxSum)
	}
	return maxSum
}

func max3(a, b, c int) int {
	return max(max(a, b), c)
}

func min3(a, b, c int) int {
	return min(min(a, b), c)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
