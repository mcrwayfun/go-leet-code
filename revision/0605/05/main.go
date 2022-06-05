package main

// time complexity: O(n)
// space complexity: O(1)
func missingNumber(nums []int) int {
	var res = len(nums)
	for i:=0; i<len(nums); i++ {
		res = res ^ i ^ nums[i]
	}
	return res
}
