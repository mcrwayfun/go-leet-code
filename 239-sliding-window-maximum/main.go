package main

import (
	"fmt"
	"math"
)

// time complexity: O(n+k)
// space complexity: O(k)
func maxSlidingWindow(nums []int, k int) []int {
	var windowStart int
	var windows []int
	var result []int

	for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
		windows = append(windows, nums[windowEnd])
		if windowEnd >= k-1 {
			result = append(result, findMaxNum(windows))
			windows = windows[1:]
			windowStart++
		}
	}

	return result
}

func findMaxNum(nums []int) int {
	if len(nums) == 0 {
		return math.MinInt64
	}
	var maxNum = nums[0]
	for _, v := range nums {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}

func main() {
	k := 3
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, k))
}

/**
题目：https://leetcode-cn.com/problems/sliding-window-maximum/

给你一个整数数组 nums，有一个大小为k的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。返回滑动窗口中的最大值。

输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

思路：使用滑动窗口的方式，维护一个窗口值和sum

解法1：
1：维护3个变量，windowStart、windowEnd、sum，窗口不停的向前滑动，并将值累加到sum。
2：当 windowEnd-windowStart > k, 开始滑动窗口：
	2-1：输出sum到result中
	2-2：滑动窗口移除先进入的值，并同时减去sum中的值
*/
