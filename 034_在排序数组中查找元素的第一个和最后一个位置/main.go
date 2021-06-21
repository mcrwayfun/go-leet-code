package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	ans := searchRange(nums, 8)
	fmt.Println(ans) // [3,4]

	ans = searchRange(nums, 6)
	fmt.Println(ans) // [-1,-1]

	ans = searchRange([]int{}, 0)
	fmt.Println(ans) // [-1,-1]

	nums = []int{5, 5, 5, 5, 8, 9}
	ans = searchRange(nums, 5)
	fmt.Println(ans) // [0,3]

	nums = []int{5, 5, 5, 5, 5}
	ans = searchRange(nums, 5)
	fmt.Println(ans) // [0,4]
}

// 时间复杂度: O(lgn)
// 空间复杂度: O(1)
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	ans := []int{-1, -1}

	// 寻找左边界
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			end = mid
		} else if nums[mid] < target {
			start = mid
		} else {
			end = mid
		}
	}

	if nums[start] == target {
		ans[0] = start
	} else if nums[end] == target {
		ans[0] = end
	} else {
		ans[0], ans[1] = -1, -1
		return ans
	}

	// 寻找右边界
	start = 0
	end = len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			start = mid
		} else if nums[mid] < target {
			start = mid
		} else {
			end = mid
		}
	}

	if nums[end] == target {
		ans[1] = end
	} else if nums[start] == target {
		ans[1] = start
	} else {
		ans[0], ans[1] = -1, -1
		return ans
	}
	return ans
}
