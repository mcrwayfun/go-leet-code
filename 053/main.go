package main

import "fmt"

func maxSubArray(nums []int) int {
	sum, ans := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if sum+nums[i] < nums[i] {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		ans = max(ans, sum)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ret := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	ans := maxSubArray(ret)
	fmt.Println(ans)
}
