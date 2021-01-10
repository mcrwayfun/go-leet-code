package main

import "fmt"

func removeElement(nums []int, val int) int {
	first, last := 0, len(nums)
	for first < last {
		if nums[first] == val {
			nums[first] = nums[last-1]
			last--
		} else {
			first++
		}
	}
	fmt.Printf("数组长度为:%d, 数据为:%v\n", len(nums), nums)
	return last
}

func main() {
	nums := []int{3, 2, 2, 3}
	target := 3
	element := removeElement(nums, target)
	fmt.Println(element)
}
