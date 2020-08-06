package main

import (
	"fmt"
)

/// Time Complexity: O(n)
/// Space Complexity: O(1)
func findMin(nums []int) int {
	min := nums[0]
	for _, v := range nums[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

/// Time Complexity: O(lgn)
/// Space Complexity: O(1)
func findMin2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return min(nums[0], nums[1])
	}

	left := 0
	right := len(nums) - 1
	for left < right {// 需要返回nums[left], 故区间需要为[left,right],此时left=right
		if nums[left] < nums[right] {
			return nums[left]// 已经处于升序的区间
		}

		mid := left + (right - left)/2
		if nums[left] <= nums[mid]  {
			left = mid + 1 // 此时[left,mid]为升序阶段,则拐点在另外半边
		}else {
			right = mid
		}
	}

	return nums[left]
}

func min(a,b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	nums := []int{3, 4, 5, 1, 2}
	min := findMin(nums)
	fmt.Println(min)
}
