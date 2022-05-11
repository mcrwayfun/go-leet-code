package main

import "fmt"

/**
题目：
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

示例 2：
输入：nums = [0,1]
输出：[[0,1],[1,0]]

示例 3：
输入：nums = [1]
输出：[[1]]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func permute(nums []int) [][]int
*/

/**
解题思路：
假设现在有数组[0,1,2]，求全排列有
0,1,2
0,2,1
1,0,2
1,2,0
2,0,1
2,1,0

假设固定下标0的元素，交换下标1和2的元素，既可以得到排列[0,1,2]和[0,2,1]
随后固定下标1的元素，交换下标0和2的元素，以此类推

time complexity: O(n*n!)
space complexity: O(1)
*/
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	var res [][]int
	permuteRec(nums, 0, &res)
	return res
}

func permuteRec(nums []int, start int, res *[][]int) {
	if start == len(nums) {
		*res = append(*res, append([]int{}, nums...)) // new 一个新的slice
	} else {
		for i := start; i < len(nums); i++ {
			nums[i], nums[start] = nums[start], nums[i]
			permuteRec(nums, start+1, res)
			nums[i], nums[start] = nums[start], nums[i]
		}
	}
}

func permuteV2(nums []int) [][]int {
	var res [][]int
	var ans []int
	backtrack(nums, ans, &res)
	return res
}

func backtrack(nums []int, ans []int, res *[][]int) {
	if len(ans) == len(nums) {
		// 注意不要直接append ans，在二维数组中，修改ans会直接影响res的值
		*res = append(*res, append([]int{}, ans...))
		return
	}

	for _, v := range nums {
		if contain(ans, v) {
			continue
		}
		ans = append(ans, v)
		backtrack(nums, ans, res)
		ans = ans[:len(ans)-1]
	}
}

func contain(ans []int, target int) bool {
	for _, v := range ans {
		if v == target {
			return true
		}
	}

	return false
}

func main() {
	ret := permuteV2([]int{5, 4, 6, 2})
	fmt.Println(ret)

	res := permute([]int{0, 1, 2})
	fmt.Println(res)
}
