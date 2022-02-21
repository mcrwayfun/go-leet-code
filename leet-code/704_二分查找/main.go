package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	ans := f2(nums, target)
	fmt.Println(ans)

	target = 2
	ans = f2(nums, target)
	fmt.Println(ans)
}

/// Time Complexity: O(lgn)
/// Space Complexity: O(1)
func f1(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			return mid
		}
	}

	return -1
}

func f2(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid
		} else {
			end = mid
		}
	}
	return -1
}
