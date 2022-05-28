package main

/**
使用递归的方式，每次固定第一个元素，然后开始交换
剩余的元素位置，达到求全排列的目标。举个例子，比如有[1,2,3]
先固定1，则有
1,2,3
1,3,2
类比则有
2,1,3
2,3,1
3,2,1
3,1,2

time complexity: O(n*n!)
space complexity: O(1)
*/
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	var res [][]int
	permuteHelper(nums, 0, &res)
	return res
}

func permuteHelper(nums []int, start int, res *[][]int) {
	if len(nums) == start {
		*res = append(*res, append([]int{}, nums...))
	} else {
		for i := start; i < len(nums); i++ {
			nums[i], nums[start] = nums[start], nums[i]
			permuteHelper(nums, start+1, res)
			nums[i], nums[start] = nums[start], nums[i]
		}
	}
}
