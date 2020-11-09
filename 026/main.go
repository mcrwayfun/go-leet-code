package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for i = 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	fmt.Printf("数组长度为:%d, 数组为:%v\n", i+1, nums[:i+1])
	return i + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	_ = removeDuplicates(nums)

	nums = []int{1, 1, 2}
	_ = removeDuplicates(nums)

	nums = []int{1, 1, 1}
	_ = removeDuplicates(nums)

	nums = []int{1}
	_ = removeDuplicates(nums)

	nums = []int{1, 1}
	_ = removeDuplicates(nums)
}
