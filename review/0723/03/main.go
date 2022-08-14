package main

import "fmt"

// time complexity: O(n*lgn)
// space complexity: O(1)
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 || len(nums) < k {
		return -1
	}

	var i = 0
	var j = len(nums) - 1
	for i <= j {
		p := partition(nums, i, j)
		if p == k-1 {
			return nums[p]
		} else if p < k-1 {
			i = p + 1
		} else {
			j = p - 1
		}
	}
	return -1
}

func partition(nums []int, low, high int) int {
	var pivot = nums[low]
	var i = low
	var j = high
	for i < j {
		for ; i < j && nums[j] < pivot; j-- {
		}
		nums[i], nums[j] = nums[j], nums[i]
		for ; i < j && nums[i] >= pivot; i++ {
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return i
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
