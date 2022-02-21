package main

import "fmt"

func subsets(nums []int) [][]int {
	var res [][]int
	var ans []int
	backtrack(0, nums, ans, &res)
	return res
}

func backtrack(index int, nums []int, ans []int, res *[][]int) {
	*res = append(*res, append([]int{}, ans...))

	for i := index; i < len(nums); i++ {
		ans = append(ans, nums[i])
		backtrack(i+1, nums, ans, res)
		ans = ans[:len(ans)-1]
	}
}


func main() {
	sub := subsets([]int{1, 2, 3})
	fmt.Println(sub)
}
