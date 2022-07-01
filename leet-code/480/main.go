package main

import "fmt"

func medianSlidingWindow(nums []int, k int) []float64 {
	var result []float64
	var windows []int

	for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
		windows = append(windows, nums[windowEnd])
		if windowEnd >= k-1 {
			result = append(result, findMedian(windows))
			windows = windows[1:]
		}
	}
	return result
}

func findMedian(nums []int) float64 {
	if len(nums)%2 != 0 {
		return float64(nums[len(nums)/2])
	}
	return float64((nums[len(nums)/2] + nums[len(nums)/2-1]) / 2)
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(medianSlidingWindow(nums, k))

	nums = []int{1, 3, -1, -3, 5, 3, 6, 7}
	k = 4
	fmt.Println(medianSlidingWindow(nums, k))
}

/**
https://leetcode-cn.com/problems/sliding-window-median/

中位数是有序序列最中间的那个数。如果序列的长度是偶数，则没有最中间的数；此时中位数是最中间的两个数的平均数。

例如：

[2,3,4]，中位数是3
[2,3]，中位数是 (2 + 3) / 2 = 2.5
给你一个数组 nums，有一个长度为 k 的窗口从最左端滑动到最右端。窗口中有 k 个数，每次窗口向右移动 1 位。你的任务是找出每次窗口移动后得到的新窗口中元素的中位数，并输出由它们组成的数组。

示例：

给出nums = [1,3,-1,-3,5,3,6,7]，以及k = 3。

窗口位置                      中位数
---------------               -----
[1  3  -1] -3  5  3  6  7       1
 1 [3  -1  -3] 5  3  6  7      -1
 1  3 [-1  -3  5] 3  6  7      -1
 1  3  -1 [-3  5  3] 6  7       3
 1  3  -1  -3 [5  3  6] 7       5
 1  3  -1  -3  5 [3  6  7]      6
因此，返回该滑动窗口的中位数数组[1,-1,-1,3,5,6]。

窗口位置                      中位数
---------------               -----
[1  3  -1 -3]  5  3  6  7       1
 1 [3  -1  -3  5] 3  6  7      -2
 1  3 [-1  -3  5  3] 6  7       1
 1  3  -1 [-3  5  3  6] 7       4
 1  3  -1  -3 [5  3  6  7]      4


解析题目：这道题给定了一个窗口值k，每当窗口滑动的时候，需要求出当前窗口的中值。当k是奇数时，中值取窗口的中间数即可。
当k为偶数，需要取最中间的两个数的平均数。

你可以假设 k 始终有效，即：k 始终小于等于输入的非空数组的元素个数。
与真实值误差在 10 ^ -5 以内的答案将被视作正确答案。

解法1：
1：初始化3个变量，windowStart表示当前窗口的初始下标，windows存储当前窗口的值，result用于存储中位数
2：遍历数组，向windows中插入值，当 windowEnd >= k-1，此时需要移动窗口。先求出当前窗口的中位数
并添加到result中，然后移除windows最早进去的元素，最后将windowStart++
*/
