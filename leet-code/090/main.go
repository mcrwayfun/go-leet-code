package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	var ans [][]int
	var res []int
	sort.Ints(nums) // nlgn
	backtrack(0, nums, res, &ans)
	return ans
}

func backtrack(start int, nums, res []int, ans *[][]int) {
	*ans = append(*ans, append([]int{}, res...))

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue // 存在重复的元素, 需要剪枝
		}

		res = append(res, nums[i])
		backtrack(i+1, nums, res, ans)
		res = res[:len(res)-1]
	}
}

func main() {
	nums := []int{1, 2, 2}
	dup := subsetsWithDup(nums)
	fmt.Println(dup)
}
