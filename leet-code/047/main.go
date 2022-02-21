package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	var ans [][]int
	var res []int
	used := make([]bool, len(nums))
	sort.Ints(nums)
	backtrack(nums, res, used, &ans)
	return ans
}

func backtrack(nums, res []int, used []bool, ans *[][]int) {
	if len(res) == len(nums) {
		*ans = append(*ans, append([]int{}, res...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if !used[i] { // 当前元素没有使用过
			//1 nums[i-1] 没用过 说明回溯到了同一层 此时接着用num[i] 则会与 同层用num[i-1] 重复
			//2 nums[i-1] 用过了 说明此时在num[i-1]的下一层 相等不会重复
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			res = append(res, nums[i])
			used[i] = true
			backtrack(nums, res, used, ans)
			used[i] = false
			res = res[:len(res)-1]
		}
	}
	return
}

func main() {
	nums := []int{1, 2, 2}
	ans := permuteUnique(nums)
	fmt.Println(ans)
}
