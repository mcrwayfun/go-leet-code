package main

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var res []int
	backtrack(0, 0, target, candidates, res, &ans)
	return ans
}

func backtrack(start, sum, target int, candidates, res []int, ans *[][]int) {
	if sum == target {
		*ans = append(*ans, append([]int{}, res...))
		return
	}

	for i:=start; i<len(candidates); i++{
		if sum > target {// 剪枝
			continue
		}
		res = append(res, candidates[i])
		backtrack(i, sum+candidates[i], target, candidates, res, ans)
		res = res[:len(res)-1]
	}
	return
}
