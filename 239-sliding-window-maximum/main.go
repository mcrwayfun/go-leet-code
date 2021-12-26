package main

import (
	"fmt"
	"math"
)

// time complexity: O(n*k)
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

// time complexity: O(n)
// space complexity: O(k)
func maxSlidingWindow2(nums []int, k int) []int {
	var windowStart int
	var result []int
	var queue = &MyQueue{}

	for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
		queue.push(nums[windowEnd])
		if windowEnd >= k-1 {
			result = append(result, queue.front())
			queue.pop(nums[windowStart])
			windowStart++
		}
	}

	return result
}

type MyQueue struct {
	data []int
}

func (q *MyQueue) isEmpty() bool {
	return len(q.data) == 0
}

func (q *MyQueue) front() int {
	return q.data[0]
}

func (q *MyQueue) back() int {
	return q.data[len(q.data)-1]
}

func (q *MyQueue) push(val int) {
	for !q.isEmpty() && val > q.back() {
		q.data = q.data[:len(q.data)-1] // remove the last element
	}
	q.data = append(q.data, val)
}

func (q *MyQueue) pop(val int) {
	if q.isEmpty() || val != q.front() {
		return
	}
	q.data = q.data[1:]
}

func main() {
	k := 3
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, k))

	k = 3
	nums = []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow2(nums, k))
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

解法1：这道题可以使用滑动窗口的思路，通过维护一个窗口，不断的向窗口中添加元素，每当窗口的长度超过阈值
k时，获取当前窗口的最大值，然后移除掉最早进入窗口的元素。重复该操作直到完成数组的遍历。

1：维护3个变量，windowStart用于标识最早进入窗口的元素的下标；windows是用于存储当前窗口中的元素；
数组result则用于存储最大值；
2：不断向滑动窗口中添加元素，当窗口中的元素超过k，则需要滑动窗口，条件是: windowEnd > k+1。这里有
为什么不可以使用 (windowEnd-windowStart)+1 > k 来判断是否已经要滑动窗口呢？因为容易漏掉边界值。
假设数组为 [1  3  -1] -3  5  3  6  7，当windowEnd为3时，此时windowStart为0，3-0+1 > 3，
需要将窗口向前移动，由于此时已经向windows中添加了下位为3的元素，所以求最大值会不准。
3：由于在滑动窗口的过程中，每一次都需要求最大值，所以时间复杂度为达到 O(n*k)，然后然后它在
leet code上超时了...

解法2：同样是滑动窗口的思想，基于解法1的前提，在滑动窗口的过程中，需要O(k)的时间去找到windows中的
最大值，这样会导致超时，需要在这里将复杂度降到O(1)

1：维护一个单调队列，队列是单调递增/递减的，队首元素存放的是最大值，需要实现以下三个方法：
	a：push(val int)
	b：pop(val int) 将指定的元素从队列中移除
	c：front() 返回队首元素
2：由于队列需要维护单调递增，且队首是最大的元素。所以push每次都需要将 < val 的元素移除，以确保
队列单调的特性（不需要将元素留在队列中，所以不需要通过移动队列中的元素来达到排序的目的）
3：pop 方法是将指定的元素从队列中移除，该方法主要是滑动窗口时，需要将窗口中最早的元素从窗口中剔除，
所以需要指定元素。如果指定的元素不是队列中的最大值(front)，也不需要将其移除。

这里解释下为什么第2种方法的时间复杂度为O(n)，明明push操作有for循环呢？
可以看下，对于数组中的任意一个元素，它被操作的次数仅为1次，所以可以简单的认为是O(1)
*/
