package main

/**
1：任何数和自己异或会得到0
2：任何数和0异或会得到自己
数组中其他数都是成对出现，将数组全部元素异或可以得到只出现一次的元素
*/
func singleNumber(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
