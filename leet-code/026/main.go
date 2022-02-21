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
	// fmt.Printf("数组长度为:%d, 数组为:%v\n", i+1, nums[:i+1])
	return i + 1
}

func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var first, last int
	for first < len(nums)-1 {
		for nums[first] == nums[last] {
			last++
			if last == len(nums) {
				return first + 1
			}
		}
		nums[first+1] = nums[last]
		first++
	}

	fmt.Printf("长度为:%d, 移除后的数组为:%v\n", len(nums), nums)
	return first + 1
}

func removeDuplicates3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}

	fmt.Printf("长度为:%d, 移除后的数组为:%v\n", last + 1, nums)
	return last + 1
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
