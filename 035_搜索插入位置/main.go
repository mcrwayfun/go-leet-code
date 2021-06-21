package main

import "fmt"

// Time Complexity: O(lgn)
// Space Complexity: O(1)
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	if nums[0] > target {
		return 0
	}

	// 找到最后一个小于target值的下标
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

	if nums[end] == target {
		return end
	}
	if nums[end] < target {
		return end + 1
	}
	if nums[start] == target {
		return start
	}
	return start + 1
}

func searchInsertV2(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	if nums[0] > target {
		return 0
	}

	// 找到第一个大于target值的下标
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

	if nums[start] >= target {
		return start
	} else if nums[end] >= target {
		return end
	} else {
		return end + 1
	}
}

func main() {
	nums := []int{1, 3, 5, 6}
	ans := searchInsert(nums, 5)
	fmt.Println(ans) // 2

	ans = searchInsert(nums, 2)
	fmt.Println(ans) // 1

	ans = searchInsert(nums, 7)
	fmt.Println(ans) // 4

	ans = searchInsertV2(nums, 0)
	fmt.Println(ans) // 0

	ans = searchInsertV2(nums, 5)
	fmt.Println(ans) // 2

	ans = searchInsertV2(nums, 2)
	fmt.Println(ans) // 1

	ans = searchInsertV2(nums, 7)
	fmt.Println(ans) // 4

	ans = searchInsertV2(nums, 0)
	fmt.Println(ans) // 0
}
