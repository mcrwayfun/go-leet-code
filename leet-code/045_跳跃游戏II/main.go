package main

import (
	"fmt"
	"math"
)

// time complexity: O(n^2)
// space complexity: O(n)
func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	steps := make([]int, len(nums))
	steps[0] = 0
	for i := 1; i < len(steps); i++ {
		steps[i] = math.MaxInt64
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if steps[j] != math.MaxInt64 && nums[j]+j >= i {
				steps[i] = min(steps[i], steps[j]+1)
			}
		}
	}

	return steps[len(nums)-1]
}

// time complexity: O(n)
// space complexity: O(1)
func jumpV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var step int
	var start, end int
	for end < len(nums)-1 {
		step++
		most := end
		for i := start; i <= end; i++ {
			if nums[i]+i > most {
				most = nums[i] + i
			}
		}
		start = end + 1
		end = most
	}

	return step
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	step := jump(nums)
	fmt.Println(step)

	nums = []int{2, 3, 1, 1, 4}
	step = jumpV2(nums)
	fmt.Println(step)
}
