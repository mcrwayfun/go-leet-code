package main

import "sort"

/**
给定一个候选人编号的集合candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。

candidates中的每个数字在每个组合中只能使用一次。

注意：解集不能包含重复的组合。

示例1:
输入: candidates =[10,1,2,7,6,1,5], target =8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]

示例2:
输入: candidates =[2,5,2,1,2], target =5,
输出:
[
[1,2,2],
[5]
]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/combination-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func combinationSum2(candidates []int, target int) [][]int {}
*/

// time complexity:O(2^n)
// space complexity: O(n^2)
func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	sort.Ints(candidates)
	var res [][]int
	var ans []int
	helper(candidates, target, 0, 0, &res, &ans)
	return res
}

func helper(candidates []int, target int, sum int, start int, res *[][]int, ans *[]int) {
	if sum == target {
		*res = append(*res, append([]int{}, *ans...))
	} else {
		for i := start; i < len(candidates); i++ {
			if sum > target {
				continue
			}
			*ans = append(*ans, candidates[i])
			helper(candidates, target, sum+candidates[i], i+1, res, ans)
			*ans = (*ans)[:len(*ans)-1]
		}
	}
}
