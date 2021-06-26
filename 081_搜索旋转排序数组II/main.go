package main

import "fmt"

func search(nums []int, target int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}

	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target := 0
	index := search(nums, target)
	fmt.Println(index)

	target = 3
	index = search(nums, target)
	fmt.Println(index)
}
