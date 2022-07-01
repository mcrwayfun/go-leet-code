package main

import "fmt"

// time complexity: O(n^2)
// space complexity: O(n)
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	lst := make([]int, len(nums))
	for i := range lst {
		lst[i] = 1 // 初始化每个木桩的步数为1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				lst[i] = max(lst[i], lst[j]+1)
			}
		}
	}

	var best int
	for i := 0; i < len(lst); i++ {
		best = max(best, lst[i])
	}

	return best
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{5, 4, 1, 2, 3}
	lis := lengthOfLIS(nums)
	fmt.Println(lis)
}
