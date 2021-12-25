package main

import "fmt"

type MovingAverage struct {
	windowSize int
	sum        int
	nums       []int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{windowSize: size}
}

// time complexity: O(1)
// space complexity: O(n)
func (this *MovingAverage) Next(val int) float64 {
	this.sum += val
	this.nums = append(this.nums, val)
	if len(this.nums) > this.windowSize {
		this.sum -= this.nums[0]
		this.nums = this.nums[1:]
	}
	return float64(this.sum) / float64(len(this.nums))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
func main() {
	obj := Constructor(3)

	next := obj.Next(1)
	fmt.Println(next)

	next = obj.Next(10)
	fmt.Println(next)

	next = obj.Next(3)
	fmt.Println(next)

	next = obj.Next(5)
	fmt.Println(next)
}

/**
题目：https://leetcode-cn.com/problems/qIsx9U/

输入：
inputs = ["MovingAverage", "next", "next", "next", "next"]
inputs = [[3], [1], [10], [3], [5]]
输出：
[null, 1.0, 5.5, 4.66667, 6.0]

解释：
MovingAverage movingAverage = new MovingAverage(3);
movingAverage.next(1); // 返回 1.0 = 1 / 1
movingAverage.next(10); // 返回 5.5 = (1 + 10) / 2
movingAverage.next(3); // 返回 4.66667 = (1 + 10 + 3) / 3
movingAverage.next(5); // 返回 6.0 = (10 + 3 + 5) / 3

解析：起始设置一个数组长度，不断地向数组中插入值，每次插入都需要求出一个平均值：数组元素和/数组长度。
当数组中的值超过数组长度时，需要移除最早进入到数组中的值，使数组中的元素始终维持在初始化的数组长度。

解法1：
1：维护一个数组和滑动窗口值，当数组长度大于滑动窗口时，将最近进入到数组中的元素移除

解法2：
1：使用一个队列，利用队列先进先出的思想（与方法1相似）
*/
