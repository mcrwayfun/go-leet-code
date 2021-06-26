package main

import "fmt"

func findMin(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	start := 0
	end := len(nums) - 1
	for start+1 < end {
		if nums[start] < nums[end] {
			return nums[start]
		}

		mid := start + (end-start)/2
		if nums[mid] > nums[end] {
			start = mid
		} else if nums[mid] < nums[end] {
			end = mid
		} else {
			end = end - 1
		}
	}

	return min(nums[start], nums[end])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	nums := []int{2, 2, 2, 0, 1}
	min := findMin(nums)
	fmt.Println(min)

	nums = []int{1, 1, 1, 1}
	min = findMin(nums)
	fmt.Println(min)
}
