package main

import "fmt"

/**
题目：Given an array of positive numbers and a positive number ‘k’,
find the maximum sum of any contiguous subarray of size ‘k’.

解析：给定一个正数数组和正数k，求任何大小为k的连续子数组的最大和

举例：
Example 1:

Input: [2, 1, 5, 1, 3, 2], k=3
Output: 9
Explanation: Subarray with maximum sum is [5, 1, 3].
Example 2:

Input: [2, 3, 4, 1, 5], k=2
Output: 7
Explanation: Subarray with maximum sum is [3, 4].

模板：
func findMaxSumSubArray(k int, arr []int) int
*/

/**
思路1：使用暴力破解法，使用双层循环的方式
1：外层循环遍历整个数组
2：内层循环遍历长度为k
3：使用全局变量 maxSum 来记录最大值

time complexity: O(n^2)
space complexity: O(1)
*/
func findMaxSumSubArray(k int, arr []int) int {
	var maxSum int

	for i := 0; i <= len(arr)-k; i++ {
		var sum int
		for j := i; j < i+k; j++ {
			sum += arr[j]
		}
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

/**
思路2：使用滑动窗口的方式
1：使用 windowStart 记录当前窗口开始值，使用全局变量 maxSum 来记录最大值
2：循环遍历整个数组 arr，使用全局变量 sum 记录当前子数组的和
3：当 windowEnd - windowStart == k 时
	3-1：比较当前是否存在最大值
	3-2：将 arr[windowStart] 从sum中移除，因为这个是即将离开窗口中的元素
	3-3：windowStart++

time complexity: O(n)
space complexity: O(1)
*/
func findMaxSumSubArray2(k int, arr []int) int {
	var windowStart int
	var maxSum int
	var sum int

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		sum += arr[windowEnd]
		if windowEnd-windowStart+1 == k {
			maxSum = max(maxSum, sum)
			sum -= arr[windowStart]
			windowStart++
		}
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// input: arr=[2, 1, 5, 1, 3, 2], k=3
	// output: 9
	arr := []int{2, 1, 5, 1, 3, 2}
	k := 3

	// input: arr=[2, 3, 4, 1, 5], k=2
	// output: 7
	arr2 := []int{2, 3, 4, 1, 5}
	k2 := 2

	maxSum := findMaxSumSubArray(k, arr)
	fmt.Println("findMaxSumSubArray, result=", maxSum)

	maxSum2 := findMaxSumSubArray(k2, arr2)
	fmt.Println("findMaxSumSubArray, result=", maxSum2)

	maxSum = findMaxSumSubArray2(k, arr)
	fmt.Println("findMaxSumSubArray2, result=", maxSum)

	maxSum2 = findMaxSumSubArray2(k2, arr2)
	fmt.Println("findMaxSumSubArray2, result=", maxSum2)
}
