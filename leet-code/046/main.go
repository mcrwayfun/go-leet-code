package main

import "fmt"

func permute(nums []int) [][]int {
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
	ret := permute([]int{5, 4, 6, 2})
	fmt.Println(ret)
}
