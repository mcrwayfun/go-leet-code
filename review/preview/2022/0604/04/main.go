package main

// time complexity: O(n!)
// space complexity: O(1)
func permute(nums []int) [][]int {
	var ans [][]int
	if len(nums) == 0 {
		return ans
	}

	permuteHelper(nums, 0, &ans)
	return ans
}

func permuteHelper(nums []int, start int, res *[][]int) {
	if start == len(nums) {
		*res = append(*res, append([]int{}, nums...))
	} else {
		for i:=start;i<len(nums);i++{
			nums[i], nums[start] = nums[start], nums[i]
			permuteHelper(nums, start+1, res)
			nums[i], nums[start] = nums[start], nums[i]
		}
	}
}
