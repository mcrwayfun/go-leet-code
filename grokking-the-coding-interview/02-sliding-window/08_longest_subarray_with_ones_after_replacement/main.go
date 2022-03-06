package main

import "fmt"

/**
题目：Given an array containing 0s and 1s, if you are allowed to replace no more than ‘k’ 0s with 1s,
find the length of the longest contiguous subarray having all 1s.

解析：给定一个数组，只包含0和1。允许你用1替换不超过k个0，请找出包含所有1的最长的连续子数组的长度。

举例：
Example 1:

Input: Array=[0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1], k=2
Output: 6
Explanation: Replace the '0' at index 5 and 8 to have the longest contiguous subarray of 1s having length 6.
Example 2:

Input: Array=[0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1], k=3
Output: 9
Explanation: Replace the '0' at index 6, 9, and 10 to have the longest contiguous subarray of 1s having length 9.

模板：
func findLength(arr []int, k int) int
*/

/**
解题思路：使用滑动窗口的方法来解决问题
1：维护一个全局变量 numOneCount, 表示为当前窗口中1的个数
2：遍历数组arr，当前元素num==1，则numOneCount++
3：如果符合窗口滑动条件，windowEnd-windowStart+1-numOneCount>k，表示当前窗口可替换的元素已经达到了k个，
则此时需要滑动窗口。
	if num == 1, numOneCount-- and windowStart++
4: 滑动窗口的过程中，其实是要找到某个窗口中包含最多的1，即numOneCount最大

time complexity: O(n)
space complexity: O(1)
*/
func findLength(arr []int, k int) int {
	if len(arr) == 0 {
		return 0
	}

	var windowStart int
	var maxLength int
	var numOneCount int

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		if arr[windowEnd] == 1 {
			numOneCount++
		}

		if windowEnd-windowStart+1-numOneCount > k {
			if arr[windowStart] == 1 {
				numOneCount--
			}
			windowStart++
		}

		maxLength = max(maxLength, windowEnd-windowStart+1)
	}

	return maxLength
}

func main() {
	arr := []int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1}
	k := 2

	arr2 := []int{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1}
	k2 := 3

	fmt.Println(findLength(arr, k))
	fmt.Println(findLength(arr2, k2))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
