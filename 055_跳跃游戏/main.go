package main

import "fmt"

// time complexity: O(n^2)
// space complexity: O(n)
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	can := make([]bool, len(nums))
	can[0] = true
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if can[j] && j+nums[j] >= i {
				can[i] = true
				break
			}
		}
	}
	return can[len(can)-1]
}

// time complexity: O(n)
// space complexity: O(n)
func canJumpV2(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	var k int
	for i := 0; i < len(nums); i++ {
		if k < i {
			return false
		}

		k = max(k, i+nums[i])
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	jump := canJump(nums)
	fmt.Println(jump)

	nums = []int{3, 2, 1, 0, 4}
	jump = canJump(nums)
	fmt.Println(jump)

	nums = []int{2, 3, 1, 1, 4}
	jump = canJumpV2(nums)
	fmt.Println(jump)

	nums = []int{3, 2, 1, 0, 4}
	jump = canJumpV2(nums)
	fmt.Println(jump)
}
