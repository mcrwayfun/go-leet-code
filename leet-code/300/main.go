package main

/**
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

示例 1：
输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。

示例 2：
输入：nums = [0,1,0,3,2,3]
输出：4

示例 3：
输入：nums = [7,7,7,7,7,7,7]
输出：1

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-increasing-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func lengthOfLIS(nums []int) int {}
*/

/**
解法一：d(i) 表示第i个数字结尾的最长递增子序列的长度
1：需要拿a(i)和之前每个数字做比较，如果a(i)>a(j), 则可以将其加到递增子序列的后面；和当前
d(i)相比，获取一个最大值。
2：d(i)和最大值相比，获取一个最大值

举个例子，当前子序列为 1, 8, 2, 6, 4, 5。
*/
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var n = len(nums)
	var maxLIS = 1
	var d = make([]int, n)
	d[0] = 1

	for i := 1; i < n; i++ {
		// nums[i] 和 [0,i) 区间中的每个数比较，得到当前最大的连续子序列
		for j := 0; j < i; j++ {
			var cur int
			// nums[i] 可以追加到 nums[j]的后面，cur=d[j]+1
			if nums[i] > nums[j] {
				cur = d[j] + 1
			} else { // 否则等于它本身，即1
				cur = 1
			}
			// 每次都需要求max d(i)
			d[i] = max(d[i], cur)
		}
		maxLIS = max(maxLIS, d[i])
	}
	return maxLIS
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
