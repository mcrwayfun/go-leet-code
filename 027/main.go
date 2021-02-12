package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
				j++
			} else {
				j++
			}
		}
	}
	return j
}

func main() {
	nums := []int{2, 3, 3, 3}
	target := 3
	element := removeElement(nums, target)
	fmt.Println(element)
}
