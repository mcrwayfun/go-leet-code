package main

// nums=0,1,3  result=3
// 3^0^0
// 3^1^1
// 3^3^2
// time complexity: O(n)
// space complexity: O(1)
func missingNumber(nums []int) int {
	var result = len(nums)
	for i := 0; i < len(nums); i++ {
		result = i ^ nums[i] ^ result
	}
	return result
}
