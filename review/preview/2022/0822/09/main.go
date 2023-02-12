package main

// time complexity: O(n)
// space complexity: O(1)
func singleNumber(nums []int) int {
	var res int
	for _, num := range nums {
		res ^= num
	}
	return res
}
