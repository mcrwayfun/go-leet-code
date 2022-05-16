package main

/**
题目：
给定一个包含 [0, n]中n个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。

示例 1：
输入：nums = [3,0,1]
输出：2
解释：n = 3，因为有 3 个数字，所以所有的数字都在范围 [0,3] 内。2 是丢失的数字，因为它没有出现在 nums 中。

示例 2：
输入：nums = [0,1]
输出：2
解释：n = 2，因为有 2 个数字，所以所有的数字都在范围 [0,2] 内。2 是丢失的数字，因为它没有出现在 nums 中。

示例 3：
输入：nums = [9,6,4,2,3,5,7,0,1]
输出：8
解释：n = 9，因为有 9 个数字，所以所有的数字都在范围 [0,9] 内。8 是丢失的数字，因为它没有出现在 nums 中。

示例 4：
输入：nums = [0]
输出：1
解释：n = 1，因为有 1 个数字，所以所有的数字都在范围 [0,1] 内。1 是丢失的数字，因为它没有出现在 nums 中。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/missing-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func missingNumber(nums []int) int
*/

/**
解法：这道题是要求[0,n]缺少的数字，可以使用异或的方式来解决问题
1：首先将数组中的元素异或一次
2：再从0-n异或一次
3：因为数字若存在则异或为0，最终得到的结果就是缺少的数字

time complexity: O(n)
space complexity: O(1)
*/
func missingNumber(nums []int) int {
	var n = len(nums)
	result := n
	for i := 0; i < len(nums); i++ {
		result = nums[i] ^ i ^ result
	}
	return result
}
