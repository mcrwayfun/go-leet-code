package main

// time complexity: O(n)
// space complexity: O(1)
func missingNumber(nums []int) int {
	var n = len(nums)
	var res = n
	for i := 0; i < n; i++ {
		res = res ^ i ^ nums[i]
	}
	return res
}
