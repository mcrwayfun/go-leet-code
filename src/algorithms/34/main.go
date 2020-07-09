package main

import "fmt"

/// Time Complexity: O(n)
/// Space Complexity: O(h)
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

func main() {
	ints := []int{5, 7, 7, 8, 8, 10}
	ret := searchRange(ints, 8)
	fmt.Println(ret)
}
