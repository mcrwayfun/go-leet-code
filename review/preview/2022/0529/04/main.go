package main

/**
利用任何数和自己异或都为0，任何数和0异或都是自己的原理。

time complexity: O(n)
space complexity: O(1)
*/
func missingNumber(nums []int) int {
	var res = len(nums)
	for i := 0; i < len(nums); i++ {
		res = res ^ nums[i] ^ i
	}
	return res
}
