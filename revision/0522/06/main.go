package main

/**
这道题考虑位运算
1：任何数和自己异或会得到0，即 a^a=0
2：任何数和0异或会得到自己，即 a^0=a

time complexity: O(n)
space complexity: O(1)
*/
func singleNumber(nums []int) int {
	var ret = 0
	for i := 0; i < len(nums); i++ {
		ret ^= nums[i]
	}
	return ret
}
