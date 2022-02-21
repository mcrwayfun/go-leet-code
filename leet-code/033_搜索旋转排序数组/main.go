package main

import "fmt"

func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}
		if nums[start] <= nums[mid] {
			if nums[start] <= target &&
				nums[mid] >= target {
				end = mid
			} else {
				start = mid
			}
		} else {
			if nums[mid] <= target &&
				nums[end] >= target {
				start = mid
			} else {
				end = mid
			}
		}
	}

	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	index := search(nums, target)
	fmt.Println(index)

	target = 3
	index = search(nums, target)
	fmt.Println(index)
}
