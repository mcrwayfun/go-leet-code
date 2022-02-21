package main

func majorityElement(nums []int) int {
	ans, count := nums[0], 1
	for i:=1;i<len(nums);i++{
		if count == 0 {
			ans, count = nums[i], 1
		} else {
			if ans == nums[i] {
				count++
			} else {
				count--
			}
		}
	}
	return ans
}
