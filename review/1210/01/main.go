package main

// time complexity: O(m)
// space complexity: O(1)
func hammingDistance(x int, y int) int {
	// 可以使用 异或 运算符得到两个数之间的不同
	// 于是这道题转换为求 异或 结果中有多少个 1
	// 求多少个1，可以使用 n&n-1, 这样只需要循环 n 的位数即可
	return numOfOne(x ^ y)
}

func numOfOne(x int) int {
	var count int
	for x != 0 {
		count++
		x = x & (x - 1)
	}
	return count
}
