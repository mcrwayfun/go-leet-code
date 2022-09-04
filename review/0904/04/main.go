package main

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	var res [][]int
	helper(nums, 0, &res)
	return res
}

func helper(nums []int, start int, res *[][]int) {
	if start == len(nums) {
		*res = append(*res, append([]int{}, nums...))
	} else {
		for i := start; i < len(nums); i++ {
			nums[i], nums[start] = nums[start], nums[i]
			helper(nums, start+1, res)
			nums[i], nums[start] = nums[start], nums[i]
		}
	}
}
