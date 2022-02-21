package main

import "fmt"

func findPeakElement(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	start := 0
	end := len(nums) - 1 // 取值防止溢出
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] < nums[mid+1] {
			start = mid
		} else {
			end = mid
		}
	}

	if nums[start] > nums[end] {
		return start
	}
	return end
}

func main() {
	peeks := []int{1, 2, 1, 3, 5, 6, 4}
	peek := findPeakElement(peeks)
	fmt.Println(peek)
}
