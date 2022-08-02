package main

/**
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，
这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，
如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

示例1：
输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。

示例 2：
输入：nums = [1,2,3,1]
输出：4
解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
    偷窃到的最高金额 = 1 + 3 = 4 。

示例 3：
输入：nums = [1,2,3]
输出：3

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/house-robber-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func rob(nums []int) int {}
*/

// 头尾相邻的各抢一次，取最大
// time complexity: O(n)
// space complexity: O(1)
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	var n = len(nums)
	return max(helper(nums, 0, n-2), helper(nums, 1, n-1))
}

func helper(nums []int, start int, end int) int {
	var prev2 int
	var prev1 int
	for i := start; i <= end; i++ {
		cur := max(prev1, prev2+nums[i])
		prev2 = prev1
		prev1 = cur
	}
	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}