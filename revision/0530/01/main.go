package main

/**
利用异或的特性
1：自己和自己异或得到的是0
2：自己和0异或得到的是自己
这道题有出的不好的地方就是，为什么res要初始化为n而不是0？
看题目给的例子都默认有0了，而且像是照着答案给例子，因为下标从0开始。

time complexity: O(n)
space complexity: O(1)
 */
func missingNumber(nums []int) int {
	var res = len(nums)
	for i := 0; i < len(nums); i++ {
		res ^= i ^ nums[i]
	}
	return res
}
