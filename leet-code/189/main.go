package main

import "fmt"

// time complexity: O(n)
// space complexity: O(1)
func rotate(nums []int, k int) {
	if nums == nil || len(nums) == 0 || k == 0 {
		return
	}

	k =  k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
	return
}

func reverse(nums []int, start, end int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	if start >= end {
		return
	}

	for ; start < end; {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
	return
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	target := 3
	rotate(nums, target)
	fmt.Println(nums)
}
