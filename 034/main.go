package main

import "fmt"

/// Time Complexity: O(n)
/// Space Complexity: O(1)
func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else { // find the nums[nums] == target
			res := []int{mid, mid}
			for res[0]-1 >= 0 && nums[res[0]-1] == target {
				res[0]--
			}

			for res[1]+1 <= len(nums)-1 && nums[res[1]+1] == target {
				res[1]++
			}

			return res
		}
	}

	return []int{-1, -1}
}

/// Time Complexity: O(n)
/// Space Complexity: O(1)
func searchRange2(nums []int, target int) []int {
	return []int{searchFirstEqualElement(nums, target), searchLastEqualElement(nums, target)}
}

func searchFirstEqualElement(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else { // nums[mid] == target
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			right = mid - 1
		}
	}

	return -1
}

func searchLastEqualElement(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else { // nums[mid] == target
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			left = mid + 1
		}
	}

	return -1
}

func searchRange3(nums []int, target int) []int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] >= target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		}
	}

	if start > len(nums)-1 || nums[start] != target {
		return []int{-1, -1}
	}
	first := start // 开始位置

	start, end = 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] <= target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		}
	}

	if end < 0 || nums[end] != target {
		return []int{-1, -1}
	}

	last := end
	return []int{first, last}
}

/**
[5,7,7,8,8,10]
8
[5,7,7,8,8,10]
6
[]
0
*/
func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	ans := f1(nums, 8)
	fmt.Println(ans)
}

func f1(nums []int, target int) []int {
	left, right := 0, len(nums)
	for left < right { // 左闭右开
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if left > len(nums)-1 || nums[left] != target {
		return []int{-1, -1}
	}
	start := left

	left, right = 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if right < 0 || nums[right] != target {
		return []int{-1, -1}
	}
	end := right
	return []int{start, end}
}
